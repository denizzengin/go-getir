package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchDbQuery : Mongo connection method
func SearchDbQuery(filter MongoHandlerRequest) ([]Record, error) {
	ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_ATLAS_URL")))
	defer client.Disconnect(ctx)
	defer c()
	if err != nil {
		return nil, ErrMongoDbConnection
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
		return nil, ErrMongoDbConnection
	}

	db := client.Database(MongoDbName)
	collection := db.Collection(MongoDbCollectionName)
	if collection == nil {
		return nil, ErrMongoDbConnection
	}

	// Filter fields parse date format
	startDate, _ := time.Parse(DateTemplate, filter.StartDate)
	endDate, _ := time.Parse(DateTemplate, filter.EndDate)
	cursor, cError := collection.Find(ctx, bson.M{"createdAt": bson.M{"$gte": startDate, "$lt": endDate}, "totalCount": bson.M{"$gte": filter.MinCount, "$lt": filter.MaxCount}})
	if cError != nil {
		return nil, ErrMongoDbConnection
	}

	var records []Record
	if err = cursor.All(ctx, &records); err != nil {
		log.Fatal(err)
		return nil, ErrMongoDbConnection
	}

	return records, nil
}

// Record response
type Record struct {
	Key        string    `bson:"key,omitempty"`
	CreatedAt  time.Time `bson:"createdAt,omitempty"`
	TotalCount int       `bson:"totalCount,omitempty"`
}
