package types

import (
	"github.com/artheus/restlib"
	"github.com/artheus/restlib/serializer"
)

type XML struct {
	contentTypeOverride string
}

const (
	ContentTypeXML = "application/xml"
)

func (m *XML) OverrideContentType(contentType string) {
	m.contentTypeOverride = contentType
}

func (m *XML) ContentType() string {
	if m.contentTypeOverride != "" { return m.contentTypeOverride }

	return ContentTypeXML
}

func (m *XML) Serializer() restlib.Serializer {
	return serializer.XML()
}
