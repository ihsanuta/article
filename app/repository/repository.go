package repository

import (
	"article/app/repository/articles"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	Articles articles.Articles
}

func Init(db *gorm.DB, rd *redis.Client) *Repository {
	repo := &Repository{
		Articles: articles.InitArticlesRepository(
			db,
			rd,
		),
	}
	return repo
}
