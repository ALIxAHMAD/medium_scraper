package scraper

import (
	"context"
	"errors"
	"fmt"
	"medium_scraper/domain/entity"
	"strings"

	"github.com/chromedp/chromedp"
)

var (
	ErrUnexpected = errors.New("unexpected error")
)

type repo struct {
}

func NewRepo() repo {
	return repo{}
}

func (r *repo) SearchArticles(SearchText string) ([]entity.Article, error) {
	var titles []string
	var intros []string
	var foundArticles []entity.Article

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(
		ctx,
		scrape(&titles, &intros, SearchText),
	)
	if err != nil {
		return nil, err
	}

	if len(titles) < 10 || len(intros) < 10 {
		return nil, ErrUnexpected
	}

	for i := 0; i < 10; i++ {
		var article = entity.Article{
			Title: titles[i],
			Intro: intros[i],
		}
		foundArticles = append(foundArticles, article)
	}
	return foundArticles, nil
}

func scrape(titles *[]string, intros *[]string, search string) chromedp.Tasks {
	search = strings.ReplaceAll(search, " ", "+")
	searchMedium := strings.Join([]string{"https://medium.com/search?q=", search}, "")
	fmt.Println(searchMedium)

	return chromedp.Tasks{
		chromedp.Navigate(searchMedium),
		chromedp.Evaluate(titleJsText, &titles),
		chromedp.Evaluate(introsJsText, &intros),
	}
}

var (
	titleJsText string = `
	function getText(){
	  var docs = document.querySelectorAll('h2')
	  var filteredDocs = []
	  for (let index = 0; index < 10; index++) {
		filteredDocs.push(docs[index].textContent)
	  }
	return filteredDocs
	};
	var a = getText(); a;
	`

	introsJsText string = `
	function getText(){
	  var docs = document.querySelectorAll('div > div > a > div > p')
	  var filteredDocs = []
      for (let index = 0; index < 10; index++) {
      filteredDocs.push(docs[index].textContent)
      }
	return filteredDocs
    };
	var a = getText(); a;
	`
)
