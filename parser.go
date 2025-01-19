package vandale

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseSearchResults(n *html.Node) ([]Entry, error) {
	var f func(*html.Node) ([]Entry, error)
	f = func(n *html.Node) ([]Entry, error) {
		if isElement(n, "div", "snippets") {
			if n.FirstChild == nil {
				return nil, errors.New("no results found")
			}
			return parseEntries(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			records, err := f(c)
			if err != nil {
				return nil, err
			}
			if records != nil {
				return records, nil
			}
		}
		return nil, nil
	}
	return f(n)
}

func isEntryContainer(n *html.Node) bool {
	return isElement(n, "span", "f0j")
}

func parseEntries(n *html.Node) ([]Entry, error) {
	var entry []Entry
	var f func(*html.Node)
	f = func(n *html.Node) {
		if isEntryContainer(n) {
			word := extractWord(n)
			recordType := extractType(n)
			meanings := extractMeanings(n)
			entry = append(entry, Entry{
				SourceWord: SourceWord{
					Word: word,
					Type: recordType,
				},
				Meanings: meanings,
			})
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return entry, nil
}

func isElement(n *html.Node, tag string, class string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	if n.Data != tag {
		return false
	}
	for _, a := range n.Attr {
		if a.Key == "class" && a.Val == class {
			return true
		}
	}
	return false
}

func isMeaningList(n *html.Node) bool {
	return isElement(n, "span", "f3 f0g")
}

func getMeaningContainer(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isElement(c, "span", "f0") {
			return c
		}
	}
	return nil
}

func isVariantExamplesSeparator(n *html.Node) bool {
	if isElement(n, "span", "fr") {
		return n.FirstChild.Data == ": "
	}
	return false
}

func isNoteOpener(n *html.Node) bool {
	if isElement(n, "span", "fq") {
		return n.FirstChild.Data == "("
	}
	return false
}

func isNoteCloser(n *html.Node) bool {
	if isElement(n, "span", "fq") {
		return n.FirstChild.Data == ") " || n.FirstChild.Data == ")"
	}
	return false
}

func isVariantSeparator(n *html.Node) bool {
	if isElement(n, "span", "fr") {
		return n.FirstChild.Data == ", "
	}
	return false
}

func isExampleSeparator(n *html.Node) bool {
	if isElement(n, "span", "fr") {
		return n.FirstChild.Data == "; "
	}
	return false
}

func isExampleLanguageSeparator(n *html.Node) bool {
	if isElement(n, "span", "fq") {
		return strings.TrimSpace(n.FirstChild.Data) == ""
	}
	return false
}

func extractMeaning(n *html.Node) Meaning {
	meaning := Meaning{
		Variants: make([]Variant, 1),
	}
	inVariants := true
	variantNumber := 0
	inWord := true
	inNote := false
	inSourceLang := true
	exampleNumber := 0
	start := n.FirstChild
	for c := start; c != nil; c = c.NextSibling {
		if inVariants {
			if inWord {
				if isNoteOpener(c) {
					inWord = false
					inNote = true
					continue
				}
				if isVariantSeparator(c) {
					meaning.Variants = append(meaning.Variants, Variant{})
					variantNumber++
					continue
				}
				if isVariantExamplesSeparator(c) {
					inVariants = false
					meaning.Examples = append(meaning.Examples, Example{})
					continue
				}
				meaning.Variants[variantNumber].Word += c.FirstChild.Data
			}
			if inNote {
				if isNoteCloser(c) {
					inNote = false
					inWord = true
					continue
				}
				meaning.Variants[variantNumber].Note += c.FirstChild.Data
			}
		}
		if !inVariants {
			if inSourceLang {
				if isExampleLanguageSeparator(c) {
					inSourceLang = false
					continue
				}
				meaning.Examples[exampleNumber].InSourceLang += c.FirstChild.Data
			}
			if !inSourceLang {
				if isExampleSeparator(c) {
					inSourceLang = true
					meaning.Examples = append(meaning.Examples, Example{})
					exampleNumber++
					continue
				}
				meaning.Examples[exampleNumber].Translation += c.FirstChild.Data
			}
		}
	}
	return meaning
}

func extractMeanings(n *html.Node) []Meaning {
	var meanings []Meaning
	var f func(*html.Node)
	f = func(n *html.Node) {
		if isMeaningList(n) {
			meanings = append(meanings, extractMeaning(getMeaningContainer(n)))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return meanings
}

func extractType(n *html.Node) string {
	var recordType string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "f3 f0g" {
					return
				}
				if a.Key == "class" && a.Val == "fq" {
					recordType = fmt.Sprintf("%s%s", recordType, n.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return recordType
}

func extractWord(n *html.Node) string {
	var word string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "span" {
			for _, a := range c.Attr {
				if a.Key != "class" {
					continue
				}
				if a.Val == "ff" || a.Val == "f0i" {
					word = fmt.Sprintf("%s%s", word, c.FirstChild.Data)
				}
				if a.Val == "fr" {
					break
				}
			}
		}
	}
	return word
}
