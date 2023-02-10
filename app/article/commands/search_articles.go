package commands

import (
	"fmt"
	"medium_scraper/domain/entity"
	"medium_scraper/domain/repository"
	"strings"
)

type SearchArticlesCommand struct {
	SearchText string
}

type SearchArticlesCommandResult struct {
	Articles []entity.SearchArticle
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
	command.SearchText = strings.ReplaceAll(command.SearchText, " ", "+")
	searchMedium := strings.Join([]string{"https://medium.com/search?q=", command.SearchText}, "")
	fmt.Println("command SearchArticles from URL: ", searchMedium)
	articles, err := h.repo.SearchArticles(searchMedium)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	result := SearchArticlesCommandResult{
		Articles: articles,
	}
	return &result, nil
}
