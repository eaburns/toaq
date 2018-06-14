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
	| egrep -v 'PrettyPrint should have comm'\
	| egrep -v '^ast/.*underscores in Go names'\
	| egrep -v '^ast/.*const __ should be _'\
	| egrep -v '^ast/.*(Start|End|PrettyPrint).*should have comment'\
	| egrep -v '^logic/.*(ASTNode|Write|Visit).*should have comment'\
	| egrep -v '^vendor/'\
	> $o 2>&1
# Silly: diff the grepped golint output with empty.
# If it's non-empty, error, otherwise succeed.
e=$(mktemp tmp.XXXXXXXXXX)
touch $e
diff $o $e > /dev/null || { rm $e; fail; }

rm $o $e
