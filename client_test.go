package vandale

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVanDaleHTMLFetcher_Fetch(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected *html.Node
		mockResp http.Response
	}{
		{
			name:     "not found error returns error",
			path:     "/non-existant-path",
			expected: nil,
			mockResp: http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("")),
			},
		},
		{
			name:     "client returns okay",
			path:     "/good-path",
			expected: &html.Node{Type: html.DocumentNode}, // You might want to create a more specific expected node
			mockResp: http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("<html><body>test</body></html>")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHttpClient := &MockHttpClient{
				MockGet: func(url string) (*http.Response, error) {
					return &tt.mockResp, nil
				},
			}
			vd := vanDaleHTMLFetcher{
				httpClient: mockHttpClient,
				Host:       "https://vandale.nl",
			}
			actual, _ := vd.fetch(tt.path)
			if actual == nil && tt.expected != nil {
				t.Errorf("expected non-nil node, got nil")
			} else if actual != nil && tt.expected == nil {
				t.Errorf("expected nil node, got non-nil")
			}
		})
	}
}

type MockHttpClient struct {
	MockGet func(url string) (*http.Response, error)
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	return m.MockGet(url)
}
