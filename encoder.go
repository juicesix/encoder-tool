package encoder_tool

// Encoder represents a format encoder
type Encoder interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
	String() string
}
