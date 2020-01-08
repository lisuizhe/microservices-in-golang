package main

import (
	"context"
	"sync"

	pb "github.com/lisuizhe/microservices-in-golang/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// InmemoryRepository ...
type InmemoryRepository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create ...
func (repo *InmemoryRepository) Create(consignment *pb.Consignment) error {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return nil
}

// GetAll ...
func (repo *InmemoryRepository) GetAll() ([]*pb.Consignment, error) {
	return repo.consignments, nil
}

// MongoRepository ...
type MongoRepository struct {
	collection *mongo.Collection
}

// Create ...
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll ...
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cur.Next((context.Background())) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
