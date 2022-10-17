package metadataparser

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/UsingCoding/fpgo/pkg/maybe"
	"github.com/pkg/errors"
	"golang.org/x/net/html"

	"telegrambot/pkg/pocket/app"
)

const (
	openGraphTitleProperty = "og:title"
	openGraphImageProperty = "og:image"

	contentAttribute = "content"
)

func NewOpenGraphParser() app.MetadataParser {
	return &openGraphParser{}
}

type openGraphParser struct{}

func (parser *openGraphParser) Parse(u *url.URL) (app.Metadata, error) {
	// Request the HTML page
	res, err := http.Get(u.String())
	if err != nil {
		return app.Metadata{}, errors.Wrap(app.ErrURLFailedToProceed, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return app.Metadata{}, errors.WithStack(app.ErrURLFailedToProceed)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return app.Metadata{}, errors.WithStack(app.ErrFailedToParseForMetadata)
	}

	var title maybe.Maybe[string]
	var imageURL *url.URL
	var crawlerErr error

	const metaTagsSelector = "head meta"
	doc.Find(metaTagsSelector).Each(func(i int, selection *goquery.Selection) {
		if crawlerErr != nil {
			return
		}

		attrs := selection.Nodes[0].Attr
		if attrs[0].Val == openGraphTitleProperty {
			content := findContentAttribute(attrs)
			if maybe.Valid(content) {
				title = content
				return
			}
		}

		if attrs[0].Val == openGraphImageProperty {
			content := findContentAttribute(attrs)
			if maybe.Valid(content) {
				imageURL, crawlerErr = url.ParseRequestURI(maybe.Just(content))
				if crawlerErr != nil {
					return
				}
			}
		}
	})
	if crawlerErr != nil {
		return app.Metadata{}, errors.Wrap(app.ErrFailedToParseForMetadata, crawlerErr.Error())
	}

	// If open graph title empty, use title tag
	if !maybe.Valid(title) {
		const titleSelector = "title"
		doc.Find(titleSelector).Each(func(i int, selection *goquery.Selection) {
			title = maybe.NewJust(selection.Text())
		})
	}

	var image maybe.Maybe[*url.URL]
	if imageURL != nil {
		image = maybe.NewJust(imageURL)

		if !maybe.Just(image).IsAbs() {
			absImageURL, err2 := url.Parse(fmt.Sprintf(
				"%s%s%s",
				u.Scheme,
				u.Host,
				imageURL.Path,
			))
			if err2 != nil {
				return app.Metadata{}, errors.WithMessage(err2, "failed to make image url absolute")
			}
			image = maybe.NewJust(absImageURL)
		}
	}
	return app.Metadata{
		Title:    title,
		ImageURL: image,
	}, nil
}

func findContentAttribute(attributes []html.Attribute) maybe.Maybe[string] {
	for _, attribute := range attributes {
		if attribute.Key == contentAttribute {
			return maybe.NewJust(attribute.Val)
		}
	}
	return maybe.NewNone[string]()
}
