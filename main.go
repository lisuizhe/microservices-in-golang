package main

import (
	"sync"

	pb "github.com/lisuizhe/microservices-in-golang/proto/consignment"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository ...
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}
