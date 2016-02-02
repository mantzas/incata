package incata

// Writer Interface for writing events to storage
type Writer interface {
	Write(Event) error
}
