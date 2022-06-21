package zinc

import (
	"context"

	"github.com/fabmation-gmbh/zinc-go/pkg/meta"
)

// IndexCreateService provides methods to manage indicies.
type IndexCreateService struct {
	cli *Client

	req struct {
		Name        string                `json:"name"`
		StorageType meta.IndexStorageType `json:"storage_type"`
		Mappings    struct {
			Props map[string]meta.IndexMappingProperty `json:"properties"`
		} `json:"mappings,omitempty"`

		// TODO: Support settings
		// Settings    *IndexSettings         `json:"settings,omitempty"`
	}
}

// NewIndexCreateService returns a fully initialized IndexCreateService.
func NewIndexCreateService(cli *Client) *IndexCreateService {
	svc := &IndexCreateService{
		cli: cli,
	}

	svc.req.Mappings.Props = make(map[string]meta.IndexMappingProperty)

	return svc
}

// Name sets the name of the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexCreateService) Name(name string) *IndexCreateService {
	i.req.Name = name

	return i
}

// IndexStorageType sets the storage type of the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexCreateService) IndexStorageType(storage meta.IndexStorageType) *IndexCreateService {
	i.req.StorageType = storage

	return i
}

// AddMappingProperty adds the given index property to the index.
//
// The IndexService object is returned to allow stacking method calls.
func (i *IndexCreateService) AddMappingProperty(prop meta.IndexMappingProperty) *IndexCreateService {
	i.req.Mappings.Props[prop.Name] = prop

	return i
}

// Create creates the index.
func (i *IndexCreateService) Create(ctx context.Context) (IndexCreateResponse, error) {
	var resp IndexCreateResponse

	_, err := i.cli.c.R().
		SetResult(&resp).
		SetBody(i.req).
		Post("/index")
	if err != nil {
		return IndexCreateResponse{}, err
	}

	return resp, nil
}

// IndexGetMappingResponse holds the index mapping.
type IndexGetMappingResponse struct {
	// Name is the index name.
	Name string `json:"name"`
	// Mappings holds the index mappings.
	Mappings meta.Mappings `json:"mappings"`
}
