package scraper

import (
	"context"
	"errors"
	"medium_scraper/domain/entity"

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

func (r *repo) SearchArticles(SearchText string) ([]entity.SearchArticle, error) {
	var titles []string
	var intros []string
	var urls []string
	var foundArticles []entity.SearchArticle

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(
		ctx,
		scrapeArticles(&titles, &intros, &urls, SearchText),
	)
	if err != nil {
		return nil, err
	}

	if len(titles) < 10 || len(intros) < 10 || len(urls) < 10 {
		return nil, ErrUnexpected
	}

	for i := 0; i < 10; i++ {
		var article = entity.SearchArticle{
			Title: titles[i],
			Intro: intros[i],
			Url:   urls[i],
		}
		foundArticles = append(foundArticles, article)
	}
	return foundArticles, nil
}

func (r *repo) GetArticle(url string) (*entity.Article, error) {
	var article entity.Article
	var text []string

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(
		ctx,
		scrapeArticle(&text, url),
	)
	if err != nil {
		return nil, err
	}

	article.Text = text
	return &article, nil
}

func scrapeArticle(text *[]string, articleUrl string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(articleUrl),
		chromedp.Evaluate(articleJsText, text),
	}
}

func scrapeArticles(titles *[]string, intros *[]string, urls *[]string, searchUrl string) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Navigate(searchUrl),
		chromedp.Evaluate(titleJsText, &titles),
		chromedp.Evaluate(introsJsText, &intros),
		chromedp.Evaluate(urlsJsText, &urls),
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

	urlsJsText string = `
	function getText(){
		var docs = document.querySelectorAll('article div > div > div > a[aria-label]')
		var filteredDocs = []
		for (let index = 0; index < docs.length; index++) {
		  filteredDocs.push(docs[index].getAttribute("href"))
		}
	   function removeDuplicates(arr) {
		  return arr.filter((item,
			  index) => arr.indexOf(item) === index);
	  }
   
	  return removeDuplicates(filteredDocs);
	  };
	  var a = getText(); a;
	`

	articleJsText string = `
	function getText() {
	  var docs = document.querySelectorAll('section div div *[data-selectable-paragraph]')
      var filteredDocs = []
      for (let index = 0; index < docs.length ; index++) {
      filteredDocs.push(docs[index].textContent)
      }
	return filteredDocs
	};
	var a = getText(); a;
	`
)
