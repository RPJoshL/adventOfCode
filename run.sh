#!/bin/sh

nodemon --quiet -e go,html,yaml --ignore web/app/ --signal SIGTERM --exec 'clear && go run ./cmd/adventOfCode/ '$1' '$2' '$3' || exit 1'