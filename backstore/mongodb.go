package backstore

import (
	"context"
	"time"

	"github.com/shitpostingio/autopostingbot/documentstore/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const opDeadline = 10 * time.Second

func FindRandomPost(collection *mongo.Collection) (post entities.Post, err error) {

	ctx, cancelCtx := context.WithTimeout(context.Background(), opDeadline)
	defer cancelCtx()

	//
	filter := bson.D{
		{
			Key: "$sample",
			Value: bson.D{
				{Key: "size", Value: 1},
			},
		},
		{
			Key: "$match",
			Value: bson.D{
				{Key: "has_error", Value: false},
				{Key: "deleted_at", Value: nil},
				{Key: "posted_at", Value: bson.E{Key: "$exists", Value: true}},
			},
		},
	}
	
	//
	cursor, err := collection.Aggregate(ctx, filter, options.Aggregate())
	if err != nil {
		return post, err
	}

	//
	err = cursor.Decode(&post)
	return post, err

}