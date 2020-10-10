package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Random will return a random meme
func (c *Client) Random(memeType MediaType) (Response, error) {
	var meme Response
	var client http.Client

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.address, "random"), nil)

	resp, err := client.Do(request)
	if err != nil {
		return meme, fmt.Errorf("get request failed: %w", err)
	}

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&meme)
	if err != nil {
		return meme, fmt.Errorf("can't open request body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return meme, fmt.Errorf("can't open request body: %w", err)
	}

	return meme, nil
}
