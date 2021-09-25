#!/bin/sh

if [ "${APP_DEBUG}" = 'true' ]; then
	CompileDaemon -build="go build -o main" -command="./main serve"
else
	go build -o main -v && chmod u+x main
	exec ./main serve
fi
