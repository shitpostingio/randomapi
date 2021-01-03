package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Random will return a random post
func (c *Client) Random(postType MediaType) (Response, error) {
	var post Response
	var client http.Client

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.address, "random"), nil)

	resp, err := client.Do(request)
	if err != nil {
		return post, fmt.Errorf("get request failed: %w", err)
	}

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&post)
	if err != nil {
		return post, fmt.Errorf("can't open request body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return post, fmt.Errorf("can't open request body: %w", err)
	}
	return post, nil
}
