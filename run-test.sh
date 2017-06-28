#!/bin/sh

cd handlers
go test -cover $(glide novendor)

exit $?
