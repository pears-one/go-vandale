package vandale

import "testing"

func TestParseLanguagesFromMode(t *testing.T) {
	tests := []struct {
		mode    string
		lang1   string
		lang2   string
		wantErr bool
	}{
		{"nl-en", "nederlands", "engels", false},
		{"en-nl", "engels", "nederlands", false},
		{"nl-fr", "nederlands", "frans", false},
		{"fr-nl", "frans", "nederlands", false},
		{"nl-unknown", "", "", true},
		{"unknown-nl", "", "", true},
		{"nl-", "", "", true},
		{"-nl", "", "", true},
		{"nl", "", "", true},
		{"", "", "", true},
	}
	for _, tt := range tests {
		lang1, lang2, err := parseLanguagesFromMode(tt.mode)
		if (err != nil) != tt.wantErr {
			t.Errorf("extractLanguages(%q) error = %v, wantErr %v",
				tt.mode, err, tt.wantErr)
		}
		if lang1 != tt.lang1 || lang2 != tt.lang2 {
			t.Errorf("extractLanguages(%q) = %q, %q, want %q, %q",
				tt.mode, lang1, lang2, tt.lang1, tt.lang2)
		}
	}
}

func TestBuildPath(t *testing.T) {
	tests := []struct {
		lang1   string
		lang2   string
		word    string
		want    string
		wantErr bool
	}{
		{"nl", "en", "mooi", "gratis-woordenboek/nl-en/vertaling/mooi", false},
		{"nl", "en", "", "gratis-woordenboek/nl-en/vertaling/", false},
		{"nl", "en", " ", "gratis-woordenboek/nl-en/vertaling/ ", false},
	}
	for _, tt := range tests {
		got, err := buildPath(tt.lang1, tt.lang2, tt.word)
		if (err != nil) != tt.wantErr {
			t.Errorf("buildPath(%q, %q, %q) error = %v, wantErr %v",
				tt.lang1, tt.lang2, tt.word, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("buildPath(%q, %q, %q) = %q, want %q",
				tt.lang1, tt.lang2, tt.word, got, tt.want)
		}
	}
}
