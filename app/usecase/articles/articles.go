//go:generate mockery --name=Articles
package articles

import (
	"article/app/models"
	repoArticles "article/app/repository/articles"
	"time"
)

type Articles interface {
	GetArticles(param models.ArticleParam) ([]models.Article, error)
	CreateArticles(form models.Article) (models.Article, error)
}

type articles struct {
	repoArticles repoArticles.Articles
}

func InitArticles(repoArticles repoArticles.Articles) Articles {
	return &articles{
		repoArticles: repoArticles,
	}
}

func (a *articles) GetArticles(param models.ArticleParam) ([]models.Article, error) {
	return a.repoArticles.GetArticles(param)
}

func (a *articles) CreateArticles(form models.Article) (models.Article, error) {
	form.Created = time.Now()
	return a.repoArticles.CreateArticles(form)
}
