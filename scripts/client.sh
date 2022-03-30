#!/usr/bin/env bash
go build -o bin/tokenclient src/client/*.go

bin/tokenclient -create -hostname localhost -port 9090 -id foo
bin/tokenclient -write -hostname localhost -port 9090 -id foo -name bar -low 0 -mid 10 -high 20
bin/tokenclient -read -hostname localhost -port 9090 -id foo
bin/tokenclient -drop -hostname localhost -port 9090 -id foo
