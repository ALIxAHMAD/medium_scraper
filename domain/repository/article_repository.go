package repository

import "medium_scraper/domain/entity"

type ArticleRepository interface {
	SearchArticles(SearchText string) ([]entity.SearchArticle, error)
	GetArticle(url string) (*entity.Article, error)
}
