#!/bin/bash

echo "building hotreload"
go build -o hotreload ./cmd/hotreload

echo "starting hotreload"

./hotreload \
--root . \
--build "go build -o ./bin/server ./testserver" \
--exec "./bin/server"
