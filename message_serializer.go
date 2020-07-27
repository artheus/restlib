package restlib

type Serializer interface {
	// Marshal will convert the message in to a byte slice for sending in request
	Marshal(*Message) ([]byte, error)

	// Unmarshal will convert a raw message into the provided pointer to a Message interfaced struct
	Unmarshal([]byte, *Message) error
}