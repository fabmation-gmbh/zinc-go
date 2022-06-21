package meta

// IndexStorageType is a storage type of an index.
type IndexStorageType uint8

//go:generate enumer -type=IndexStorageType -json -trimprefix=IndexStorage -transform=lower

const (
	// IndexStorageUnknown means that the storage type of the index is unkown.
	IndexStorageUnknown IndexStorageType = iota
	// IndexStorageS3 is the S3 storage backend of the index.
	IndexStorageS3
	// IndexStorageMinIO is the MinIO storage backend of the index.
	IndexStorageMinIO
	// IndexStorageDisk is the Disk storage backend of the index.
	IndexStorageDisk
)

// IndexPropertyType is the type of an index property.
type IndexPropertyType int16

//go:generate enumer -type=IndexPropertyType -json -trimprefix=IdxProperty -transform=lower

const (
	// IdxPropertyUnknown means that the given property type is unknown or not supported by this library.
	IdxPropertyUnknown IndexPropertyType = iota
	// IdxPropertyText means the property is simple text.
	IdxPropertyText
	// IdxPropertyKeyword means the property is a keyword.
	// The value should usually not contain whitespaces.
	IdxPropertyKeyword
	// IdxPropertyDate means the property represents a date.
	// The default format is 2006-01-02T15:04:05Z07:00 (time.RFC3339 in GoLang).
	//
	// TODO: Support alias: time, datetime
	IdxPropertyDate
	// IdxPropertyNumeric means the property is a numeric type.
	//
	// TODO: Support alias: integer, double, long, short, int, float
	IdxPropertyNumeric
	// IdxPropertyBoolean means the property is a boolean.
	//
	// TODO: Support alias: bool
	IdxPropertyBoolean
)

// IndexMappingProperty is a mapping property of an index.
type IndexMappingProperty struct {
	// Name is the Name of the property.
	Name string `json:"-"`
	// FieldType is the type of the property.
	FieldType IndexPropertyType `json:"type"`
	// Analyzer is the used analyzer of the field.
	Analyzer Analyzer `json:"analyzer,omitempty"`
	// SearchAnalyzer is NOT SUPPORTED!
	// TODO: Support this field
	SearchAnalyzer string `json:"search_analyzer,omitempty"`
	// Format holds the field Format.
	// For date it is yyyy-MM-dd HH:mm:ss || yyyy-MM-dd || epoch_millis.
	Format string `json:"format,omitempty"`
	// Index describes whether the field should be indexed or not.
	Index bool `json:"index"`
	// Store describes whether the original value shall be stored.
	Store bool `json:"store"`
	// storable describes whether the field is storable.
	Sortable bool `json:"sortable"`
	// Aggregatable describes whether aggregation is enabled for this field.
	Aggregatable bool `json:"aggregatable"`
	// Highlightable describes whether the "highlight" feature can be used with this field.
	Highlightable bool `json:"highlightable"`
}

// NewIndexMappingProperty returns a new IndexMappingProperty object.
func NewIndexMappingProperty(name string, fieldType IndexPropertyType) IndexMappingProperty {
	prop := IndexMappingProperty{
		Name:          name,
		FieldType:     fieldType,
		Index:         true,
		Store:         false,
		Sortable:      false,
		Aggregatable:  false,
		Highlightable: false,
	}

	switch fieldType {
	case IdxPropertyDate:
		prop.Format = "2006-01-02T15:04:05Z07:00"
		prop.Sortable = true
		prop.Aggregatable = true
	case IdxPropertyNumeric:
		prop.Sortable = true
		prop.Aggregatable = true
	}

	return prop
}
