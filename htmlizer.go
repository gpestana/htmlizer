package htmlizer

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type Tag struct {
	Type  string
	Value string
}

type Htmlizer struct {
	Tags []Tag
}

func New() Htmlizer {
	return Htmlizer{}
}

func (h *Htmlizer) Load(s string) error {
	parsingValid := false
	currentTag := ""

	r := strings.NewReader(s)
	tz := html.NewTokenizer(r)
	for {
		tok := tz.Next()
		switch {
		case tok == html.ErrorToken:
			return nil
		case tok == html.StartTagToken && !parsingValid:
			t := tz.Token()
			if validTag(t.String()) {
				parsingValid = true
				currentTag = t.String()
			}
		case tok == html.EndTagToken && parsingValid:
			t := tz.Token()
			if validTag(t.String()) {
				parsingValid = false
				currentTag = ""
			}
		default:
			if parsingValid {
				val := string(tz.Text())
				tag := Tag{currentTag, val}
				tags := append(h.Tags, tag)
				h.Tags = tags
			}
		}
	}
	return nil
}

// Returns all values of `tagType`
func (h *Htmlizer) GetValues(tagType string) ([]Tag, error) {
	if valid := validTag(tagType); !valid {
		return nil, errors.New(fmt.Sprintf("Tag %v is not valid", tagType))
	}
	tags := []Tag{}
	for _, t := range h.Tags {
		if t.Type == tagType {
			tags = append(tags, t)
		}
	}
	return tags, nil
}

func validTag(tag string) bool {
	tagTypes := []string{
		"<a>", "<p>", "<h1>", "<h2>", "<h3>", "<h4>", "<h5>", "<h6>",
		"</a>", "</p>", "</h1>", "</h2>", "</h3>", "</h4>", "</h5>", "</h6>"}
	for _, t := range tagTypes {
		if t == tag {
			return true
		}
	}
	return false
}

func (h *Htmlizer) HumanReadable() string {
	str := []string{}
	for _, tag := range h.Tags {
		str = append(str, tag.String())
	}
	return strings.Join(str, "\n")
}

func (t *Tag) String() string {
	return t.Value
}
