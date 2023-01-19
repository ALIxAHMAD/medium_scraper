package commands

import (
	"fmt"
	"medium_scraper/domain/entity"
	"medium_scraper/domain/repository"
)

type SearchArticlesCommand struct {
	SearchText string
}

type SearchArticlesCommandResult struct {
	Articles []entity.Article
}

type SearchArticlesCommandHandler interface {
	Handle(
		command SearchArticlesCommand,
	) (
		*SearchArticlesCommandResult,
		error,
	)
}

type searchArticlesCommandHandler struct {
	repo repository.ArticleRepository
}

func NewSearchArticlesCommandHandler(
	repo repository.ArticleRepository,
) SearchArticlesCommandHandler {
	return searchArticlesCommandHandler{
		repo: repo,
	}
}

func (h searchArticlesCommandHandler) Handle(
	command SearchArticlesCommand,
) (
	*SearchArticlesCommandResult,
	error,
) {
	articles, err := h.repo.SearchArticles(command.SearchText)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	result := SearchArticlesCommandResult{
		Articles: articles,
	}
	return &result, nil
}
