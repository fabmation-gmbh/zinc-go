package zinc

import (
	"context"

	"github.com/fabmation-gmbh/zinc-go/pkg/meta"
)

// IndexService provides methods to manage indicies.
type IndexService struct {
	cli *Client

	// name is the index name.
	name        string
	storageType meta.IndexStorageType
	props       map[string]meta.IndexMappingProperty
}

// NewIndexService returns a fully initialized IndexService.
func NewIndexService(cli *Client) *IndexService {
	return &IndexService{
		cli:   cli,
		props: make(map[string]meta.IndexMappingProperty),
	}
}

// Name sets the name of the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexService) Name(name string) *IndexService {
	i.name = name

	return i
}

// IndexStorageType sets the storage type of the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexService) IndexStorageType(storage meta.IndexStorageType) *IndexService {
	i.storageType = storage

	return i
}

// AddMappingProperty adds the given index property to the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexService) AddMappingProperty(prop meta.IndexMappingProperty) *IndexService {
	i.props[prop.Name] = prop

	return i
}

// Create creates the index.
func (i *IndexService) Create(ctx context.Context) (IndexCreateResponse, error) {
	type props struct {
		Props map[string]meta.IndexMappingProperty `json:"properties"`
	}
	type idxCreateReq struct {
		Name        string `json:"name"`
		StorageType string `json:"storage_type"`
		Mappings    props  `json:"mappings,omitempty"`
	}

	data := idxCreateReq{
		Name:        i.name,
		StorageType: i.storageType.String(),
		Mappings: props{
			Props: i.props,
		},
	}

	var resp IndexCreateResponse

	_, err := i.cli.c.R().
		SetResult(&resp).
		SetBody(data).
		Post("/index")
	if err != nil {
		return IndexCreateResponse{}, err
	}

	return resp, nil
}
