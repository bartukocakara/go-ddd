package main

import (
    "fmt"
	"github.com/bartukocakara/go-project-structures/flat/scraper"
)

func main() {
    url := "https://example-news.com"  // Replace with your target site
    headlines, err := scraper.FetchHeadlines(url)
    if err != nil {
        fmt.Println("Error fetching headlines:", err)
        return
    }

    fmt.Println("Latest Headlines:")
    for i, headline := range headlines {
        fmt.Printf("%d. %s\n", i+1, headline)
    }
}
