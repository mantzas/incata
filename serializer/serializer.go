package serializer

// Serializer interface
type Serializer interface {
	Serialize(interface{}) (string, error)
}
