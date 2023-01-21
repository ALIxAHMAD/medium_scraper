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
	titlesJs := jsGetText(`0`)
	headersJs := jsGetText(`1`)

	search = strings.ReplaceAll(search, " ", "+")
	searchMedium := strings.Join([]string{"https://medium.com/search?q=", search}, "")
	fmt.Println(searchMedium)

	return chromedp.Tasks{
		chromedp.Navigate(searchMedium),
		chromedp.Evaluate(titlesJs, &titles),
		chromedp.Evaluate(headersJs, &intros),
	}
}

func jsGetText(sel string) (js string) {
	const funcJS = `function getText(sel) {
		var text = [];
		var elements = document.body.querySelectorAll("div.l.ci.jv div.l a.ad.ae.af.ag.ah.ai.aj.ak.al.am.an.ao.ap.aq.ar");

		for(var i = 0; i < elements.length; i++) {
			var current = elements[i];
			
			// Check the element has no children && that it is not empty
				text.push(current.children[sel].children[0].textContent);
			
		}
		return text
	 };`

	invokeFuncJS := `var a = getText('` + sel + `'); a;`
	return strings.Join([]string{funcJS, invokeFuncJS}, " ")
}
