#!/bin/bash
go build -o bin/tokenmanager main.go

# if [[ ! -f $HOME/.tokenmanager ]]
# then
#   tee -a $HOME/.tokenmanager > /dev/null <<EOF
# host: localhost
# port: 3456
# EOF
# fi

bin/tokenmanager client create --host localhost --port 3456 --id apple
bin/tokenmanager client write --host localhost --port 3456 --id apple --name mac --low 0 --mid 10 --high 20
bin/tokenmanager client read --host localhost --port 3456 --id apple

bin/tokenmanager client create --host localhost --port 3456 --id orange
bin/tokenmanager client write --host localhost --port 3456 --id orange --name tangy --low 5 --mid 18 --high 35
bin/tokenmanager client read --host localhost --port 3456 --id orange

