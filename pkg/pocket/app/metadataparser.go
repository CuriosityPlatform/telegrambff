package app

import (
	"net/url"

	"github.com/UsingCoding/fpgo/pkg/maybe"
)

type Metadata struct {
	Title    maybe.Maybe[string]
	ImageURL maybe.Maybe[*url.URL]
}

type MetadataParser interface {
	Parse(url *url.URL) (Metadata, error)
}
