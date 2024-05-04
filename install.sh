#!/bin/bash
## UNCOMMENT IF YOU WANT TO BUILD FROM SOURCE
rm goinspired
go build main.go 
mv main goinspired

sudo install -m733 goinspired /usr/bin/goinspired
