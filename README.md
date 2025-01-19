# Van Dale Go Client

The Van Dale Go Client is a Go package for interacting with the Van Dale dictionary. This package allows you to search for words and retrieve translations and meanings in different languages.

## Installation

To use the Van Dale Go Client in your project, you first need to install it by running:

```bash
go get github.com/pears-one/go-vandale
```

Ensure that your project is set up with Go Modules for dependency management.

## Usage

Here's a simple example on how to use the Van Dale Go Client to search for a word:

```go
package main

import (
    "fmt"
    "github.com/pears-one/go-vandale"
)

func main() {
    // Search for the translation of the word "hoi" from Dutch to German
    result, err := vandale.Search("hoi", "nl-du")
    if err != nil {
        fmt.Printf("Error searching for word: %s\n", err)
        return
    }

    // Print the search result
    fmt.Printf("Search Result: %+v\n", result)
}
```

## Contributing

Contributions to the Van Dale Go Client are welcome. Please submit a pull request or open an issue to discuss proposed changes.
