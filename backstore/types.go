package backstore

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Backstore handles storage operations for memes
type Backstore struct {
	gorm       *gorm.DB
	host       hostConfigs
	api        netConfigs
	collection *mongo.Collection
}

type hostConfigs struct {
	memeFolder        string
	symlinkMemeFolder string
}

type netConfigs struct {
	endpoint        string
	storageEndpoint string
}

//MemeDetails contains all the info that will be returned by the api
type MemeDetails struct {
	ID        string
	URL       string
	Caption   string
	Filename  string
	FilePath  string
	MessageID int
	MediaType string
	PostedAt  *time.Time
}

//Request describes one request sent to memesapi
type Request struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Type     int                `bson:",omitempty"`
	Meme     int
	Platform string
	User     string
	Time     time.Time
}
