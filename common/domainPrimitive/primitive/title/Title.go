package titlePrimitive

import (
	"strings"
	"unicode/utf8"
)

const TitleMaxLength = 240

type Title string

func TitleFrom(text string) (Title, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return "", ErrTitleIsEmpty
	}

	if utf8.RuneCountInString(text) > TitleMaxLength {
		return "", ErrTitleTooLong
	}

	return Title(text), nil
}

func (t Title) String() string {
	return string(t)
}
