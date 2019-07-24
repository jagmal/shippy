package main

import (
	"context"
	"log"
	pb "github.com/jagmal/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/jagmal/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repo repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - this creates a consignment object as passed as input
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// call the vessel service's client with consignment details to get the
	// matching vessel details back
	vesselResponse, err := h.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err = h.repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}