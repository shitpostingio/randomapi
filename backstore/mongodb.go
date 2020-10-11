package backstore

import (
	"context"
	"github.com/shitpostingio/autopostingbot/documentstore/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const opDeadline = 10 * time.Second

func FindRandomPost(collection *mongo.Collection) (post entities.Post, err error) {

	//
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
	}

	//
	result := collection.FindOne(ctx, filter, options.FindOne())
	if result.Err() != nil {
		return post, result.Err()
	}

	//
	err = result.Decode(&post)
	return post, err

}
