package client

//Client is an istance of memeaspi client
type Client struct {
	address  string
	platform string
}

//New creates an instance of client
func New(address, platform string) *Client {
	return &Client{
		address:  address,
		platform: platform,
	}
}
