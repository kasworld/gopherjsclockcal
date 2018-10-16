#!/usr/bin/env bash

rm jsclockcal.js*

echo "gopherjs build jsclockcal.go"
gopherjs build jsclockcal.go

go run dir2http.go