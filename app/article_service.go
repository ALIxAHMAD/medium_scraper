package app

import (
	"medium_scraper/app/article/commands"
	"medium_scraper/domain/repository"
)

type ArticleServiceCommands struct {
	SearchArticlesCommandHandler commands.SearchArticlesCommandHandler
}

type ArticleService struct {
	Commands ArticleServiceCommands
}

func NewArticleService(
	articleRepo repository.ArticleRepository,
) ArticleService {
	return ArticleService{
		ArticleServiceCommands{
			SearchArticlesCommandHandler: commands.NewSearchArticlesCommandHandler(articleRepo),
		},
	}
}
