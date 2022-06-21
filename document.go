package zinc

import "context"

// DocumentService provides methods to create documents.
type DocumentService struct {
	cli *Client

	index string
	data  any
}

// NewDocumentService returns a fully initialized DocumentService.
func NewDocumentService(cli *Client) *DocumentService {
	return &DocumentService{
		cli: cli,
	}
}

// SetIndex sets the target index of the document.
func (d *DocumentService) SetIndex(index string) *DocumentService {
	d.index = index

	return d
}

// SetData sets the target data of the document.
func (d *DocumentService) SetData(data any) *DocumentService {
	d.data = data

	return d
}

// Create creates a new document.
func (d *DocumentService) Create(ctx context.Context) (DocumentCreateResponse, error) {
	var resp DocumentCreateResponse

	_, err := d.cli.c.R().
		SetResult(&resp).
		SetBody(d.data).
		SetPathParam("index", d.index).
		Post("/{index}/_doc")
	if err != nil {
		return DocumentCreateResponse{}, err
	}

	return resp, nil
}

// CreateWithID creates a new document with a custom ID.
func (d *DocumentService) CreateWithID(ctx context.Context, id string) (DocumentCreateResponse, error) {
	var resp DocumentCreateResponse

	_, err := d.cli.c.R().
		SetResult(&resp).
		SetBody(d.data).
		SetPathParam("index", d.index).
		SetPathParam("id", id).
		Post("/{index}/_doc/{id}")
	if err != nil {
		return DocumentCreateResponse{}, err
	}

	return resp, nil
}
