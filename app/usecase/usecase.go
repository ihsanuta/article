package usecase

import (
	"article/app/repository"
	"article/app/usecase/articles"
)

type Usecase struct {
	Articles articles.Articles
}

func Init(repository *repository.Repository) *Usecase {
	uc := &Usecase{
		Articles: articles.InitArticles(
			repository.Articles,
		),
	}
	return uc
}
