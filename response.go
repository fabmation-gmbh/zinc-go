package zinc

import "github.com/fabmation-gmbh/zinc-go/pkg/meta"

// ResponseMessage is a basic response from the server.
type ResponseMessage string

const (
	// ResponseOk indicates that the request was successful.
	ResponseOk ResponseMessage = "ok"
)

// HTTPResponse is the generic HTTP response.
type HTTPResponse struct {
	Message ResponseMessage `json:"message,omitempty"`
	Error   string          `json:"error,omitempty"`
	ID      string          `json:"id,omitempty"`
	Index   string          `json:"index,omitempty"`
	Data    interface{}     `json:"data,omitempty"`
}

// IndexCreateResponse is the response returned by the server when creating an index.
type IndexCreateResponse struct {
	Message ResponseMessage `json:"message"`
	// Index is the name of the new index.
	Index string `json:"index"`
	// StorageType is the storage type of the new index.
	StorageType meta.IndexStorageType `json:"storage_type"`
	// Error is the error string returned by the server.
	Error string `json:"error"`
}

// IndexDeleteResponse is the response returned by the server when deleting an index.
type IndexDeleteResponse struct {
	Message ResponseMessage `json:"message"`
	// Index is the name of the deleted index.
	Index string `json:"index"`
	// StorageType is the storage type of the deleted index.
	StorageType meta.IndexStorageType `json:"storage"`
	// Error is the error string returned by the server.
	Error string `json:"error"`
}

// DocumentCreateResponse is the response returned by the server when creating a document.
type DocumentCreateResponse struct {
	Message ResponseMessage `json:"message"`
	// ID is the document ID.
	ID string `json:"id"`
	// Error is the error string returned by the server.
	Error string `json:"error"`
}
