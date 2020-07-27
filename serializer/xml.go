package serializer

import (
	"encoding/xml"
	"stash.sbab.se/TNT/ansible2structurizr/openstack/restlib"
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