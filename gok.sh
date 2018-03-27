#!/bin/sh
#
# Verifies that go code passes go fmt, go vet, golint, and go test.
#

o=$(mktemp tmp.XXXXXXXXXX)

fail() {
	echo Failed
	cat $o
	rm $o
	exit 1
}

trap fail INT TERM

echo Formatting
gofmt -l $(find . -name '*.go') > $o 2>&1
test $(wc -l $o | awk '{ print $1 }') = "0" || fail

echo Vetting
go vet ./... > $o 2>&1 || fail

echo Testing
go test -test.timeout=60s ./... > $o 2>&1 || fail

echo Linting
golint ./... \
	| egrep -v '^parser/.*underscores in Go names'\
	| egrep -v '^ast/.*(Start|End|Mod) should have comment'\
	> $o 2>&1
# Silly: diff the grepped golint output with empty.
# If it's non-empty, error, otherwise succeed.
e=$(tempfile)
touch $e
diff $o $e > /dev/null || { rm $e; fail; }

rm $o $e
