package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/shitpostingio/randomapi/rest/client"
)

// Random will return a random meme
func Random(w http.ResponseWriter, r *http.Request) {

	memeDetails, err := GetRandomMeme("")
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	m := client.Response{
		ID: memeDetails.ID,
		Meme: client.Data{
			URL:      memeDetails.URL,
			Filename: memeDetails.Filename,
			Type:     memeDetails.MediaType,
			Date:     memeDetails.PostedAt,
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

	return
}

func getRandomMeme() {

}
