package serializer

import (
	"encoding/json"
	"stash.sbab.se/TNT/ansible2structurizr/openstack/restlib"
)

type jsonSerializer struct {}

func (j *jsonSerializer) Marshal(message *restlib.Message) ([]byte, error) {
	return json.Marshal(message.Body)
}

func (j *jsonSerializer) Unmarshal(bytes []byte, message *restlib.Message) error {
	return json.Unmarshal(bytes, message.Body)
}

// JSON creates and returns an JSON message serializer
func JSON() restlib.Serializer {
	return &jsonSerializer{}
}



