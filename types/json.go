package types

import (
	"stash.sbab.se/TNT/ansible2structurizr/openstack/restlib"
	"stash.sbab.se/TNT/ansible2structurizr/openstack/restlib/serializer"
)

type JSON struct {
	contentTypeOverride string
}

const (
	ContentTypeJSON = "application/json"
)

func (m *JSON) OverrideContentType(contentType string) {
	m.contentTypeOverride = contentType
}

func (m *JSON) ContentType() string {
	if m.contentTypeOverride != "" { return m.contentTypeOverride }

	return ContentTypeJSON
}

func (m *JSON) Serializer() restlib.Serializer {
	return serializer.JSON()
}
