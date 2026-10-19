#!/bin/sh
# a.txt and b.txt each ship alone, never together: a combination only a
# merge queue's run sees.
if [ -e a.txt ] && [ -e b.txt ]; then
	echo "a.txt and b.txt cannot ship together" >&2
	exit 1
fi
echo "a.txt and b.txt apart"
