package client

//Client is an istance of randomapi client
type Client struct {
	address string
}

//New creates an instance of client
func New(address string) *Client {
	return &Client{
		address: address,
	}
}
