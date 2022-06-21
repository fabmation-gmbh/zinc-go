package zinc

import (
	"encoding/json"

	"github.com/fabmation-gmbh/zinc-go/pkg/meta"
	"github.com/go-resty/resty/v2"
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

// CreateIndex returns a new IndexCreateService object allowing to create a new index.
func (c *Client) CreateIndex() *IndexCreateService {
	return NewIndexCreateService(c)
}

// DeleteIndex deletes the given index.
func (c *Client) DeleteIndex(name string) (IndexDeleteResponse, error) {
	var resp IndexDeleteResponse

	_, err := c.c.R().
		SetResult(&resp).
		SetPathParam("index", name).
		Delete("/index/{index}")

	return resp, err
}

// ListIndexes returns a list of all indexes.
func (c *Client) ListIndexes() ([]meta.Index, error) {
	ret := make([]meta.Index, 0, 10)

	_, err := c.c.R().
		SetResult(&ret).
		Get("index")
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetIndexMapping returns the mapping of the given index.
func (c *Client) GetIndexMapping(name string) (IndexGetMappingResponse, error) {
	var resp IndexGetMappingResponse

	_, err := c.c.R().
		SetResult(&resp).
		SetPathParam("index", name).
		Get("/{index}/_mapping")

	return resp, err
}

// DocumentService returns a new DocumentService object.
func (c *Client) DocumentService() *DocumentService {
	return NewDocumentService(c)
}
