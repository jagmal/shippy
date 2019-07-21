rm .\consignment-service
protoc -I. --go_out=plugins=micro:. .\proto\consignment\consignment.proto
set GOOS=linux
set GOARCH=amd64
go build
