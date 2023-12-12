#!/bin/bash

if [ $# -eq 0 ]; then
    storage="in-memory"
else
    storage="$1"
fi

go run cmd/app/app.go --storage="$storage"
