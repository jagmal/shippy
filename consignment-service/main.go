// shippy-service-consignment/main.go
package main

import (
	"context"
	"log"
	"os"

	// Import generated protobuf code
	pb "github.com/jagmal/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/jagmal/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port = ":50051"
	defaultHost = "mongodb://datastore:27017"
)

func main() {

	// Create a new service
	srv := micro.NewService(
		// this name must match with the package name given in the protobuf definition
		micro.Name("shippy.service.consignment"),
	)
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

	h := &handler{repository, vesselClient}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
