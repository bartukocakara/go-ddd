package scraper

import (
    "errors"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

// FetchHeadlines fetches news headlines from the given URL
func FetchHeadlines(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, errors.New("Failed to fetch page")
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    var headlines []string
    doc.Find("h2.headline").Each(func(index int, element *goquery.Selection) {
        headline := element.Text()
        headlines = append(headlines, headline)
    })

    return headlines, nil
}
