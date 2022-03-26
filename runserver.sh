#!/bin/bash
go build -o bin/tokenmanager main.go

# if [[ ! -f $HOME/.tokenmanager ]]
# then
#   tee -a $HOME/.tokenmanager > /dev/null <<EOF
# host: localhost
# port: 3456
# EOF
# fi

bin/tokenmanager server --port 3456
