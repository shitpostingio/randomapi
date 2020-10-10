package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Response represent a meme.
type Response struct {
	ID    string `json:"id,omitempty"`
	Meme  Data   `json:"meme,omitempty"`
	Error string `json:"error,omitempty"`
}

// Data contains data info about a random meme
type Data struct {
	URL       string     `json:"url"`
	Caption   string     `json:"caption,omitempty"`
	Filename  string     `json:"filename"`
	MessageID int        `json:"messageid"`
	Type      string     `json:"mediatype"`
	Date      *time.Time `json:"date"`
}

// Random will return a random meme
func (c *Client) Random(memeType, start, end, userid string) (Response, error) {
	var meme Response
	var client http.Client

	var b strings.Builder
	fmt.Fprintf(&b, "/random?")

	if memeType != "" {
		fmt.Fprintf(&b, fmt.Sprintf("type=%s&", memeType))
	}

	if start != "" {
		fmt.Fprintf(&b, fmt.Sprintf("startDate=%s&", start))
	}

	if end != "" {
		fmt.Fprintf(&b, fmt.Sprintf("endDate=%s&", end))
	}

	built := b.String()

	if memeType != "" || start != "" || end != "" {
		built = built[0 : len(built)-1]
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.address, built), nil)
	request.Header.Add("X-user-platform", c.platform)
	request.Header.Add("X-user-id", userid)

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
