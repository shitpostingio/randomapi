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

func FindRandomPost(collection *mongo.Collection) (entities.Post, error) {

	ctx, cancelCtx := context.WithTimeout(context.Background(), opDeadline)
	defer cancelCtx()

	//
	pipeline := mongo.Pipeline{
		{
			{
				Key: "$match",
				Value: bson.D{
					{Key: "haserror", Value: nil},
					{Key: "deletedat", Value: nil},
					{Key: "postedat", Value: bson.D{{Key: "$exists", Value: true}}},
				},
			},
		},
		{
			{
				Key: "$sample",
				Value: bson.D{
					{Key: "size", Value: 1},
				},
			},
		},
	}

	//
	cursor, err := collection.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		return entities.Post{}, err
	}

	var posts []entities.Post
	err = cursor.All(ctx, &posts)
	if err != nil || len(posts) == 0 {
		return entities.Post{}, err
	}

	return posts[0], err

}
