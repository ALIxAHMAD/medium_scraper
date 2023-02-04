package commands

import (
	"fmt"
	"medium_scraper/domain/entity"
	"medium_scraper/domain/repository"
)

type GetArticleCommand struct {
	Url string
}

type GetArticleCommandResult struct {
	Article entity.Article
}

type GetArticleCommandHandler interface {
	Handle(
		command GetArticleCommand,
	) (
		*GetArticleCommandResult,
		error,
	)
}
type getArticleCommandHandler struct {
	repo repository.ArticleRepository
}

func NewGetArticleCommandHandler(
	repo repository.ArticleRepository,
) GetArticleCommandHandler {
	return getArticleCommandHandler{repo: repo}
}

func (h getArticleCommandHandler) Handle(
	command GetArticleCommand,
) (
	*GetArticleCommandResult,
	error,
) {
	article, err := h.repo.GetArticle(command.Url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	result := GetArticleCommandResult{
		Article: *article,
	}
	return &result, nil
}
