package app

import "medium_scraper/domain/repository"

type Services struct {
	ArticleService ArticleService
}

func NewServices(
	articleRepo repository.ArticleRepository,
) Services {
	return Services{
		ArticleService: NewArticleService(articleRepo),
	}
}
