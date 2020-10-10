package backstore

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rs/xid"
	"github.com/shitpostingio/randomapi/config"
	"github.com/shitpostingio/randomapi/entities"

	// mysql gorm adapter
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var fixed = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

// NewBackstore returns a new instance of Backstore, with default settings.
func NewBackstore(c *config.Config, collection *mongo.Collection) (*Backstore, error) {
	var b Backstore
	var err error // username, password, host, database
	b.gorm, err = gorm.Open("mysql", dbConnString(c.MariaDB.Username, c.MariaDB.Password, c.MariaDB.Host, c.MariaDB.DatabaseName))
	if err != nil {
		return nil, fmt.Errorf("cannot build Backstore due to gorm error: %w", err)
	}

	b.host = hostConfigs{
		memeFolder:        c.MemeFolder,
		symlinkMemeFolder: c.MemeSymlinkFolder,
	}

	b.api = netConfigs{
		storageEndpoint: c.StorageEndpoint,
		endpoint:        c.Endpoint,
	}

	b.collection = collection

	return &b, nil
}

//GetRandomMeme returns the local path of a random meme and its caption
func (b *Backstore) GetRandomMeme(startDate, endDate, mediaType, platform, userid string) (MemeDetails, error) {

	s, e, sErr, eErr := getDatesFromStrings(startDate, endDate)
	if (startDate != "" && sErr != nil) || (endDate != "" && eErr != nil) {
		return MemeDetails{}, fmt.Errorf("cannot parse date")
	}

	request := Request{
		Platform: platform,
		User:     userid,
	}

	var ext string
	var media entities.Type

	randQuery := b.gorm.Not("posted_at IS NULL OR message_id = 0")

	if mediaType != "" && (mediaType == "video" || mediaType == "image" || mediaType == "animation") {
		b.gorm.Where("name = ?", mediaType).First(&media)
		randQuery = randQuery.Where("type_id = ?", media.ID)

		// if a specific media type is requested, save it
		request.Type = int(media.ID)
	}

	var meme entities.Post
	if (s != time.Time{}) {
		randQuery = randQuery.Where("posted_at >= ?", s)
	}

	if (e != time.Time{}) {
		randQuery = randQuery.Where("posted_at <= ?", e)
	}

	randQuery.Order("rand()").First(&meme)

	var mtype string

	if meme.TypeID == 1 {
		ext = ".jpg"
		mtype = "image"
	} else {
		ext = ".mp4"
		if meme.TypeID == 2 {
			mtype = "video"
		} else {
			mtype = "animation"
		}
	}

	request.Meme = int(meme.ID)
	request.Time = time.Now()

	id := xid.New().String()
	filename := fmt.Sprintf("rand_%s%s", id, ext)

	newPath := fmt.Sprintf("%s%s", b.host.symlinkMemeFolder, filename)

	err := os.Symlink(b.host.memeFolder+meme.FileID+ext, newPath)
	if err != nil {
		return MemeDetails{}, fmt.Errorf("could not create file symlink: %w", err)
	}

	m := MemeDetails{
		ID:        id,
		URL:       fmt.Sprintf("%s/%s", b.api.storageEndpoint, filename),
		Caption:   meme.Caption,
		Filename:  filename,
		MessageID: meme.MessageID,
		FilePath:  newPath,
		MediaType: mtype,
		PostedAt:  meme.PostedAt,
	}

	go storeMessage(request, b.collection)

	return m, nil
}

func getDatesFromStrings(d1, d2 string) (parsed1 time.Time, parsed2 time.Time, sErr error, eErr error) {
	s, sErr := strconv.ParseInt(d1, 10, 64)
	if sErr == nil {
		parsed1 = time.Unix(s, 0)
	}

	e, eErr := strconv.ParseInt(d2, 10, 64)
	if eErr == nil {
		parsed2 = time.Unix(e, 0)
	}

	return
}

//storeMessage serializes and saves a message in the database
func storeMessage(r Request, c *mongo.Collection) {

	ctx, cancelCtx := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancelCtx()

	_, err := c.InsertOne(ctx, &r)
	if err != nil {
		log.Println(fmt.Sprintf("Error while inserting the request into the document store: %s", err))
		return
	}

}
