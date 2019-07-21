// shippy-service-consignment/main.go
package main

import (
	"context"
	"log"

	// Import generated protobuf code
	pb "github.com/jagmal/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastor
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

// Method to create a new consignment (also implementing repository interface)
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// Method to get all consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment,res *pb.Response) error {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the 'Response' message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments - 
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	// Create a new service
	srv := micro.NewService(
		// this name must match with the package name given in the protobuf definition
		micro.Name("shippy.service.consignment"),
	)
	srv.Init()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
