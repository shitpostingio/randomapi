package rest

import (
	"encoding/json"
	"net/http"
	"errors"
	"strings"

	"github.com/shitpostingio/randomapi/rest/client"
)

// Random will return a random meme
func (i *Interface) Random(w http.ResponseWriter, r *http.Request) {

	start := r.URL.Query().Get("startDate")
	end := r.URL.Query().Get("endDate")
	media := r.URL.Query().Get("type")

	platform := r.Header.Get("X-user-platform")
	userid := r.Header.Get("X-user-id")
	if platform == "" || userid == "" {
		writeError(w, errors.New("missing headers"), http.StatusBadRequest)
		return
	}

	memeDetails, err := i.backstore.GetRandomMeme(start, end, media, platform, userid)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	m := client.Response{
		ID: memeDetails.ID,
		Meme: client.Data{
			URL:       memeDetails.URL,
			Caption:   memeDetails.Caption,
			MessageID: memeDetails.MessageID,
			Filename:  memeDetails.Filename,
			Type:      memeDetails.MediaType,
			Date:      memeDetails.PostedAt,
		},
	}

	if strings.HasPrefix(r.Host, "localhost") || strings.HasPrefix(r.Host, "127.0.0.1") || strings.HasPrefix(r.Host, "::1") {
		m.Meme.URL = memeDetails.FilePath
	}

	jenc := json.NewEncoder(w)
	err = jenc.Encode(m)

	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}
}