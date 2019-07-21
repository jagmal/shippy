set GOOS=linux
set GOARCH=amd64
go build
docker build -t shippy-cli-consignment .