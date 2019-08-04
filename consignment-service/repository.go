package main

import (
	"context"
	"fmt"

	pb "github.com/jagmal/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository -
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		fmt.Println("ERROR WHILE RUNNING FIND() command")
		fmt.Println(err)
		fmt.Println("-------------")
	}
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
