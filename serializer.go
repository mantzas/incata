package incata

// Serializer interface
type Serializer interface {
	Serialize(interface{}) (interface{}, error)
}
