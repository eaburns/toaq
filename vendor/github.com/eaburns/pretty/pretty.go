// Package pretty provides a pretty-printer for Go types. It produces a
// lightweight, Go-syntax-like output. It elides some type information
// and syntactic details. The intent is to show a data structure, such
// as an abstract syntax tree, without much clutter.
package pretty

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Indent is the string used to denote a single level of indentation.
// New lines are indented by a series of Indents, based on the level of nesting.
var Indent = "\t"

// A PrettyPrinter implements the PrettyPrint method.
type PrettyPrinter interface {
	// PrettyPrint returns a string, overriding the default pretty-print format.
	PrettyPrint() string
}

// Fprint prints a pretty-looking version of a value to an io.Writer.
//
// If a type implementing PrettyPrinter is encountered then its PrettyPrint
// method is used to print it.
//
// When printing maps with keys that are strings, integer types, floating point
// types, or bools, elements are printed in increasing order of their keys (for bools,
// false < true). When printing maps with any other type of key, elements are
// printed in an arbitrary order, which may differ with each call to Fprint.
//
// Fprint prunes cycles. Recall that passing a value makes a copy. The copy is not
// part of a cycle. If this is undesired, pass a pointer to the value. See the PassPointer
// and PassValue examples.
func Fprint(out io.Writer, v interface{}) (err error) {
	defer func() {
		if r := recover(); r == nil {
			return
		} else if e, ok := r.(error); ok {
			err = e
		} else {
			panic(err)
		}
	}()
	print(out, make(map[reflect.Value]bool), "\n", reflect.ValueOf(v))
	return err
}

// Print prints a pretty-looking version of a value to os.Stdout.
func Print(v interface{}) error {
	return Fprint(os.Stdout, v)
}

// String prints a pretty-looking version of a value, returning it as a string.
func String(v interface{}) string {
	buf := bytes.NewBuffer(nil)
	if err := Fprint(buf, v); err != nil {
		panic(err)
	}
	return buf.String()
}

func print(out io.Writer, path map[reflect.Value]bool, indent string, v reflect.Value) {
	if !v.IsValid() {
		pr(out, "nil")
		return
	}
	if path[v] {
		pr(out, "<cycle>")
		return
	}
	path[v] = true
	defer func() { path[v] = false }()
	if pper, ok := v.Interface().(PrettyPrinter); ok {
		if v.Kind() == reflect.Ptr && v.IsNil() {
			pr(out, "nil")
		} else {
			pr(out, "%s", pper.PrettyPrint())
		}
		return
	}
	switch v.Kind() {
	case reflect.Bool:
		pr(out, "%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		pr(out, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		pr(out, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		pr(out, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		pr(out, "%f", v.Complex())

	case reflect.Array, reflect.Slice:
		printArray(out, path, indent, v)

	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			pr(out, "nil")
		} else {
			print(out, path, indent, v.Elem())
		}

	case reflect.String:
		pr(out, strconv.Quote(v.String()))

	case reflect.Struct:
		printStruct(out, path, indent, v)

	case reflect.Map:
		printMap(out, path, indent, v)

	case reflect.Chan:
		pr(out, "<chan>")
	case reflect.Func:
		pr(out, "<function>")
	case reflect.UnsafePointer:
		pr(out, "<unsafe pointer>")
	case reflect.Invalid:
		pr(out, "<invalid>")
	}
}

func printArray(out io.Writer, path map[reflect.Value]bool, indent string, v reflect.Value) {
	if v.Len() == 0 {
		pr(out, "[]")
		return
	}
	pr(out, "[")
	indent2 := indent + Indent
	for i := 0; i < v.Len(); i++ {
		pr(out, indent2)
		print(out, path, indent2, v.Index(i))
	}
	pr(out, indent+"]")
}

func printStruct(out io.Writer, path map[reflect.Value]bool, indent string, v reflect.Value) {
	t := v.Type()
	pr(out, "%s{", t.Name())
	indent2 := indent + Indent

	var u, e bool
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !exported(&f) {
			u = true
			continue
		}
		e = true
		pr(out, "%s%s: ", indent2, f.Name)
		print(out, path, indent2, v.Field(i))
	}
	if !e {
		// No exported fields, so don't put '}' on a new line.
		indent = ""
		indent2 = ""
	}
	if u {
		pr(out, "%s…", indent2)
	}
	pr(out, "%s}", indent)
}

func printMap(out io.Writer, path map[reflect.Value]bool, indent string, v reflect.Value) {
	t := v.Type()
	pr(out, "%s{", t.Name())
	indent2 := indent + Indent
	keys := v.MapKeys()
	sort.Sort(values(keys)) // Just a best-effort sorting.
	for _, k := range keys {
		pr(out, "%s", indent2)
		print(out, path, indent2, k)
		pr(out, ": ")
		print(out, path, indent2, v.MapIndex(k))
	}
	pr(out, "%s}", indent)
}

type values []reflect.Value

func (vs values) Len() int      { return len(vs) }
func (vs values) Swap(i, j int) { vs[i], vs[j] = vs[j], vs[i] }
func (vs values) Less(i, j int) bool {
	a, b := vs[i], vs[j]
	if a.Kind() != b.Kind() {
		panic("different kinds")
	}

	switch a.Kind() {
	case reflect.Bool:
		var aint, bint int
		if a.Bool() {
			aint = 1
		}
		if b.Bool() {
			bint = 1
		}
		return aint < bint

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() < b.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() < b.Uint()

	case reflect.Float32, reflect.Float64:
		return a.Float() < b.Float()

	case reflect.String:
		return a.String() < b.String()

	default:
		// Don't bother sorting any other kinds.
		return false
	}
}

func pr(out io.Writer, f string, args ...interface{}) {
	if _, err := fmt.Fprintf(out, f, args...); err != nil {
		panic(err)
	}
}

func exported(f *reflect.StructField) bool {
	r, _ := utf8.DecodeRuneInString(f.Name)
	return unicode.IsUpper(r)
}
