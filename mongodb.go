package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"time"
)

// Mongo connection method
func Mongo() {
	//ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDbAtlassURL))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	r, e := client.ListDatabases(context.TODO(), bson.D{{}})

	if e != nil {
		fmt.Println("yes")
	}

	for _, v := range r.Databases {
		fmt.Println(v.Name)
	}
}

// Record response
type Record struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Key        string             `bson:"key,omitempty"`
	CreatedAt  string             `bson:"createdAt,omitempty"`
	TotalCount int                `bson:"totalCount,omitempty"`
}
