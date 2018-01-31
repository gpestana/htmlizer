package htmlizer

import (
	"errors"
	"fmt"
	"strings"
)

// Valid tags that HTMLizer deals with

type Tag struct {
	Type  string
	Value string
}

type Htmlizer struct {
	Tags []*Tag
}

func New(html string) Htmlizer {
	return Htmlizer{}
}

// Returns all values of `tagType`
func (h *Htmlizer) GetValues(tagType string) ([]Tag, error) {
	if valid := isTagValid(tagType); !valid {
		return nil, errors.New(fmt.Sprintf("Tag %v is not valid", tagType))
	}
	tags := []Tag{}
	for _, t := range h.Tags {
		if t.Type == tagType {
			tags = append(tags, *t)
		}
	}
	return tags, nil
}

func isTagValid(tag string) bool {
	tagTypes := map[string]int{"a": 1, "p": 2, "h1": 3, "h2": 4, "h3": 5, "h4": 6, "h5": 7, "h6": 8}
	_, ok := tagTypes[tag]
	return ok
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
