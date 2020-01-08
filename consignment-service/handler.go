package main

import (
	"context"
	"log"

	pb "github.com/lisuizhe/microservices-in-golang/consignment-service/proto/consignment"
	vesselpb "github.com/lisuizhe/microservices-in-golang/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselpb.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselpb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	err = s.repository.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
