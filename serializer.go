package incata

// Serializer interface
type Serializer interface {
	Serialize(interface{}) (string, error)
}
