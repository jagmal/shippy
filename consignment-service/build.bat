
# This will auto generate the protobuf code required.
# Note: we should consider generating this code in the docker file while
# deployment but we need this to be pushed to github (so others can pull in the
# module). But we need a better automated way of regenerating this file every
# time the protobuf definition is changed in .proto file
protoc -I. --go_out=plugins=micro:. .\proto\consignment\consignment.proto

docker build -t shippy-service-consignment .
