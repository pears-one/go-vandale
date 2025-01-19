package vandale

import (
	"errors"
	"fmt"
	"strings"
)

var longNameByAbbreviation = map[string]string{
	"nl": "nederlands",
	"en": "engels",
	"fr": "frans",
	"du": "duits",
	// The languages below take a different form
	// on the website, thus they are not supported currently
	// "sp": "spaans",
	// "it": "italiaans",
	// "pt": "portugees",
	// "zw": "zweeds",
}

// parseLanguagesFromMode takes a string in the form of "nl-en" and returns the
// corresponding names in the long form. For example, "nl-en" is converted
// to "nederlands" and "engels". If the mode is not in the correct form or if
// the language is not available, an error is returned.
func parseLanguagesFromMode(mode string) (string, string, error) {
	msg := "mode should be in the form of 'nl-en'"
	if len(mode) != 5 {
		return "", "", errors.New(msg)
	}
	languages := strings.Split(mode, "-")
	if len(languages) != 2 {
		return "", "", errors.New(msg)
	}
	msg = "language is not available %s"
	lang1, ok := longNameByAbbreviation[languages[0]]
	if !ok {
		return "", "", fmt.Errorf(msg, languages[0])
	}
	lang2, ok := longNameByAbbreviation[languages[1]]
	if !ok {
		return "", "", fmt.Errorf(msg, languages[1])
	}
	return lang1, lang2, nil
}

// buildPath builds the path for the given languages and word
// the output path is in the form gratis-woordenboek/<lang1>-<lang2>/vertaling/<word>
func buildPath(lang1, lang2, word string) (string, error) {
	return fmt.Sprintf("gratis-woordenboek/%s-%s/vertaling/%s", lang1, lang2, word), nil
}
