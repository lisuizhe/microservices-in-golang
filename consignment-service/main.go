package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/lisuizhe/microservices-in-golang/consignment-service/proto/consignment"
	vesselpb "github.com/lisuizhe/microservices-in-golang/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.consignment.service"),
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
	vesselClient := vesselpb.NewVesselServiceClient("shippy.vessel.service", srv.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
