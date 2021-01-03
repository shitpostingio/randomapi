package backstore

import (
	"context"
	"errors"
	"time"

	"github.com/shitpostingio/autopostingbot/documentstore/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const opDeadline = 10 * time.Second

// FindRandomPost select one random post from yor mongo collection
func FindRandomPost(collection *mongo.Collection) (entities.Post, error) {

	//
	pipeline := mongo.Pipeline{
		{
			{
				Key: "$sample",
				Value: bson.D{
					{Key: "size", Value: 1},
				},
			},
		},
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
	}

	for i := 0; i < 3; i++ {

		//
		cursor, err := collection.Aggregate(context.TODO(), pipeline, options.Aggregate())
		if err != nil {
			return entities.Post{}, err
		}

		var posts []entities.Post
		err = cursor.All(context.TODO(), &posts)
		if err != nil || len(posts) == 0 {
			continue
		}

		return posts[0], err

	}

	return entities.Post{}, errors.New("no post found, try again")

}
