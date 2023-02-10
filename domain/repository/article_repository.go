package repository

import "medium_scraper/domain/entity"

type ArticleRepository interface {
	SearchArticles(SearchText string) ([]entity.SearchArticle, error)
	GetArticle(ArticleUrl string) (*entity.Article, error)
}
