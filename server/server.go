package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/eaburns/peggy/peg"
	"github.com/eaburns/pretty"
	"github.com/eaburns/toaq/ast"
	"github.com/eaburns/toaq/logic"
)

func main() {
	http.Handle("/font/", http.FileServer(http.Dir("./server")))
	http.HandleFunc("/", rootHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("./server/template")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.ExecuteTemplate(w, "template", nil); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		input, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		resp := make(map[string]interface{})

		p := ast.NewParser(string(input))
		parseTree, err := p.Tree()
		if err != nil {
			resp["Error"] = err.Error()
		} else {
			text, err := p.Text() // cannot err, since Tree() succeeded above.
			if err != nil {
				panic("impossible error: " + err.Error())
			}
			if text == nil {
				panic("impossible nil text")
			}
			resp["Braces"] = ast.BracesString(text)
			resp["AST"] = pretty.String(text)
			resp["Text"] = toString(text)
			resp["Parse"] = peg.Pretty(parseTree)
			switch st, err := safeInterp(text); {
			case err != nil:
				resp["Logic"] = err.Error()
			case st == nil:
				resp["Logic"] = "fragment"
			default:
				resp["Logic"] = logic.PrettyString(st)
			}
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	default:
		msg := "method " + req.Method + " is not allowed"
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}

}

// safeInterp calls the interpreter and recovers any panics.
// The interpreter should not panic, but currently it does
// on some things that are unimplemented.
func safeInterp(text *ast.Text) (st logic.Statement, err error) {
	defer func() {
		if str, ok := recover().(string); ok {
			err = errors.New("bug: " + str)
		}
	}()
	return logic.Interpret(text), nil
}

func toString(node ast.Node) string {
	var s strings.Builder
	ast.Visit(node, ast.FuncVisitor(func(n ast.Node) {
		if w, ok := n.(*ast.Word); ok {
			s.WriteString(w.T)
		}
	}))
	return s.String()
}
