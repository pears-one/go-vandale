package vandale

type SourceWord struct {
	Word string
	Type string
}

type SearchResult struct {
	SearchWord string
	SourceLang string
	TargetLang string
	Entries    []Entry
}

type Entry struct {
	SourceWord SourceWord
	Meanings   []Meaning
}

type Meaning struct {
	Variants []Variant
	Examples []Example
}

type Variant struct {
	Word string
	Note string
}

type Example struct {
	InSourceLang string
	Translation  string
}
