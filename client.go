package vandale

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// httpClient is an interface that was created for testing purposes, so we
// could mock the http.Client Get function
type httpClient interface {
	Get(url string) (*http.Response, error)
}

type htmlFetcher interface {
	Fetch(path string) (*html.Node, error)
}

type vanDaleHTMLFetcher struct {
	httpClient
	Host string
}

func defaultFetcher() vanDaleHTMLFetcher {
	return vanDaleHTMLFetcher{
		httpClient: http.DefaultClient,
		Host:       "https://www.vandale.nl",
	}
}

// fetch fetches the HTML page at the given path and returns the root
// element of the HTML document
func (vd vanDaleHTMLFetcher) fetch(path string) (*html.Node, error) {
	url := fmt.Sprintf("%s/%s", vd.Host, path)
	resp, err := vd.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d when fetching %s", resp.StatusCode, url)
	}
	defer resp.Body.Close()
	body, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML body of %s: %w", url, err)
	}
	return body, nil
}
