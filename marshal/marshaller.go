package marshal

// Serializer interface
type Serializer interface {
	Serialize(interface{}) (interface{}, error)
}

// Deserializer interface
type Deserializer interface {
	Deserialize(value interface{}, result interface{}) error
}
