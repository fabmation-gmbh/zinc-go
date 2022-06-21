package zinc

import "github.com/fabmation-gmbh/zinc-go/pkg/meta"

// ResponseMessage is a basic response from the server.
type ResponseMessage string

const (
	// ResponseOk indicates that the request was successful.
	ResponseOk ResponseMessage = "ok"
)

// IndexCreateResponse is the response returned by the server when creating an index.
type IndexCreateResponse struct {
	Message ResponseMessage `json:"message"`
	// Index is the name of the new index.
	Index string `json:"index"`
	// StorageType is the storage type of the new index.
	StorageType meta.IndexStorageType `json:"storage_type"`
}

// DocumentCreateResponse is the response returned by the server when creating a document.
type DocumentCreateResponse struct {
	Message ResponseMessage `json:"message"`
	// ID is the document ID.
	ID string `json:"id"`
}
