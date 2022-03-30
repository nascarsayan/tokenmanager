#!/usr/bin/env bash
go build -o bin/tokenserver src/server/*.go
bin/tokenserver -port 9090
