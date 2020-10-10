package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Request describes one request sent to memesapi
type request struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Type     int                `bson:",omitempty"`
	Meme     int
	Platform string
	User     string
	Time     time.Time
}

type netConfigs struct {
	endpoint        string
	storageEndpoint string
}

// Post represents a post in the document store.
type Post struct {

	// ID is MongoDB's object ID.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Media contains information on the media added.
	Media Media

	// HasError becomes true if there was an error while trying to
	// post the media.
	HasError bool `bson:",omitempty"`

	// PostedAt is the timestamp of the post on the channel.
	PostedAt *time.Time `bson:",omitempty"`

	// DeletedAt is the timestamp of the deletion from the channel.
	DeletedAt *time.Time `bson:",omitempty"`
}

// Media represents a media in the document store.
type Media struct {

	// Type is the media type.
	Type string

	// FileID is Telegram's ID of the media.
	// This ID is just temporary.
	FileID string
}
