#!/bin/sh
#
# Pull the latest grammar from toaq.org
# and diff it with the grammar
# of this repository (toaq_orig.peg).
#
# This script was modified from
# https://github.com/lojban/ilmentufa/blob/master/pegpoho.sh,
# which came with the following notice:
# Copyright (c) 2014 lagleki, ilmen, and other contributors.

rm -f toaqlanguage.js.peg
wget --quiet http://toaq.org/parser/toaqlanguage.js.peg

awk '
	/^$/ { next }
	/^ *\057/ { next }
	/^\173/,/^\175/ { next }
	{
	sub(/=/, "<-");
	sub(/expr:/, "");
	sub(/ +\173.*$/, "");
	sub(/pre:/, "");
	sub(/post:/, "");
	print }
	' toaqlanguage.js.peg | egrep -v '^$' > toaq_new.peg

diff -rup toaq_orig.peg toaq_new.peg
