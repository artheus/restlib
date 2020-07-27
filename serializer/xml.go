package serializer

import (
	"encoding/xml"
	"github.com/artheus/restlib"
)

type xmlSerializer struct {}

func (x *xmlSerializer) Marshal(message *restlib.Message) ([]byte, error) {
	return xml.Marshal(message.Body)
}

func (x *xmlSerializer) Unmarshal(bytes []byte, message *restlib.Message) error {
	return xml.Unmarshal(bytes, message.Body)
}

// XML creates and returns an XML message serializer
func XML() restlib.Serializer {
	return &xmlSerializer{}
}