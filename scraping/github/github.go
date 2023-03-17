package github

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

type Repository struct {
	Author string
	Name   string
	Link   string
	Desc   string
}

const githubStarUrl = "https://github.com/forever-z-133?tab=stars"

var record = map[string]struct{}{}

func Scraping() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	repos := make([]*Repository, 0)
	c.OnHTML(".d-lg-flex .col-12", func(e *colly.HTMLElement) {
		repo := &Repository{}
		text := e.ChildText("div.mb-1 > h3 > a")
		split := strings.Split(text, "/")
		repo.Author = strings.TrimSpace(split[0])
		repo.Name = strings.TrimSpace(split[1])
		repo.Link = e.Request.AbsoluteURL(e.ChildAttr("div.mb-1 > h3 > a", "href"))
		repo.Desc = e.ChildText("div.py-1 > p")
		repos = append(repos, repo)
	})

	c.OnHTML(".paginate-container .BtnGroup", func(e *colly.HTMLElement) {
		texts := e.ChildTexts("a")
		next := false
		for _, text := range texts {
			if text == "Next" {
				next = true
			}
		}
		if !next {
			return
		}
		urls := e.ChildAttrs("a", "href")
		url := ""
		if len(urls) == 2 {
			url = urls[1]
		} else if len(urls) == 1 {
			url = urls[0]
		}
		_, ok := record[url]
		if !ok {
			record[url] = struct{}{}
			fmt.Println(url)
			c.Visit(url)
		}
	})

	c.Visit(githubStarUrl)

	fmt.Println(len(repos))
	f, _ := os.Create("repo.json")
	encoder := json.NewEncoder(f)
	encoder.Encode(repos)
}
