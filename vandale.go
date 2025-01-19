package vandale

// Search takes a word and a language mode and returns a SearchResult for the
// word. The language mode should be in the form of "nl-en". The SearchResult
// contains the translation of the word, the source language and the target
// language.
func Search(word string, mode string) (SearchResult, error) {
	client := defaultFetcher()
	var err error
	sr := SearchResult{SearchWord: word}

	sr.SourceLang, sr.TargetLang, err = parseLanguagesFromMode(mode)
	if err != nil {
		return sr, err
	}

	path, err := buildPath(sr.SourceLang, sr.TargetLang, word)
	if err != nil {
		return sr, err
	}

	node, err := client.fetch(path)
	if err != nil {
		return sr, err
	}

	sr.Entries, err = parseSearchResults(node)
	return sr, err
}
