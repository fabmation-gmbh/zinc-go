package zinc

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	// "github.com/goccy/go-json"
)

// Client is a ZincSearch client.
type Client struct {
	c *resty.Client
}

// NewClient returns a fully initialized client.
func NewClient(opts ...clientOpt) *Client {
	c := &Client{}

	c.c = resty.New().
		SetHeader("Accept", "application/json")

	c.c.JSONMarshal = json.Marshal
	c.c.JSONUnmarshal = json.Unmarshal

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// IndexService returns a new IndexService object.
func (c *Client) IndexService() *IndexService {
	return NewIndexService(c)
}

// DocumentService returns a new DocumentService object.
func (c *Client) DocumentService() *DocumentService {
	return NewDocumentService(c)
}

