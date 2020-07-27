package restlib

type MessageType interface {
	// ContentType retrieves the MIME-type of the message body
	ContentType() string

	// Serializer retrieves the message serializer to use
	Serializer() Serializer
}
