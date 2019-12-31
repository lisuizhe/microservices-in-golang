package main

import (
	"fmt"
	"sync"

	"context"

	pb "github.com/lisuizhe/microservices-in-golang/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
)

// const (
// 	port = ":50051"
// )

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository ...
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create ...
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// GetAll ...
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Response{Created: true, Consignment: consignment}, nil
// }

// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
// 	consignments := s.repo.GetAll()
// 	return &pb.Response{Consignments: consignments}, nil
// }

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterShippingServiceServer(s, &service{repo})
	// reflection.Register(s)
	// log.Println("Running on port:", port)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
