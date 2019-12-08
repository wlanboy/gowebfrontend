# gowebfrontend
Golang webserver using mux and html templates

# build
go get -d -v

go clean

go build

# run
go run main.go

# debug
go get -u github.com/go-delve/delve/cmd/dlv
dlv debug ./goservice

# call
curl -X POST http://127.0.0.1:8000/event -H 'Content-Type: application/json' -d '{"Name": "test", "Type": "info"}'

curl -X GET http://127.0.0.1:8100/ 
