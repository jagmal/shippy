
# TODO: Find a way to automatically generate the go file whenever proto file
# is modified
Set-Item -Path Env:Path -Value ($Env:Path + ";" + $Env:GOPATH + "\bin")
protoc -I. --go_out=plugins=micro:. .\proto\vessel\vessel.proto



# set GOOS=linux
# set GOARCH=amd64
# set CGO_ENABLED=0
# go build -a -installsuffix cgo -o shippy-service-vessel


docker build -t shippy-service-vessel .
