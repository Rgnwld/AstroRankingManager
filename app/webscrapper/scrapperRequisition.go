package webscrapper

import (
	"errors"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type NewsObject struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

// func FetchNews() ([]NewsObject, error) {
// 	c := colly.NewCollector()

// 	var newsList []NewsObject

// 	// Find and visit all links
// 	c.OnHTML("#devlog a[href]", func(e *colly.HTMLElement) {
// 		newsObj := NewsObject{Url: e.Attr("href"), Title: e.Text}
// 		fmt.Println(fmt.Sprint(newsObj))
// 		newsList = append(newsList, newsObj)
// 	})

// 	c.Visit("https://rgnwld.itch.io/astro")

// 	if len(newsList) <= 0 {
// 		return nil, errors.New("Something went wrong")
// 	}

// 	return newsList, nil
// }

var newsList []NewsObject
var nextReq (time.Time)

func FetchNews() ([]NewsObject, error) {
	c := colly.NewCollector()

	if len(newsList) > 0 && nextReq.After(time.Now()) {
		fmt.Println("Using cached news")
		return newsList, nil
	}

	var newsListAux []NewsObject

	c.OnHTML("#devlog a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML(".primary_column", func(e *colly.HTMLElement) {
		newsObj := NewsObject{Url: e.Request.URL.String(), Title: e.ChildText(".post_header>h1"), Content: e.ChildText(".post_body")}
		fmt.Println(fmt.Sprint(newsObj))
		newsListAux = append(newsListAux, newsObj)
	})

	c.Visit("https://rgnwld.itch.io/astro")

	if len(newsListAux) <= 0 {
		return nil, errors.New("Something went wrong")
	}

	newsList = newsListAux

	nextReq = time.Now().Add(5 * time.Minute)

	return newsList, nil
}
