#!/bin/sh

cd handlers
go test -cover

exit $?
