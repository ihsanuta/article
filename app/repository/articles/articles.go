package articles

import (
	"article/app/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/snappy"
	"github.com/jinzhu/gorm"
)

const (
	ArticlesKeyByQuery                       = `articles:q:%s`
	durationArticlesExpiration time.Duration = 60 * time.Second
)

type Articles interface {
	GetArticles(param models.ArticleParam) ([]models.Article, error)
	CreateArticles(form models.Article) (models.Article, error)
}

type articles struct {
	db *gorm.DB
	rd *redis.Client
}

func InitArticlesRepository(db *gorm.DB, rd *redis.Client) Articles {
	return &articles{
		db: db,
		rd: rd,
	}
}

func (a *articles) CreateArticles(form models.Article) (models.Article, error) {
	return a.createArticleDB(form)
}

func (a *articles) GetArticles(param models.ArticleParam) ([]models.Article, error) {
	var (
		results []models.Article
	)
	// cek data redis
	results, err := a.getArticlesRedis(param)
	if err != nil && err != redis.Nil {
		return results, fmt.Errorf("get article redis error : %s", err.Error())
	}

	if len(results) < 1 {
		results, err := a.getArticlesDB(param)
		if err != nil {
			return results, fmt.Errorf("get article redis error : %s", err.Error())
		}
		_, err = a.setArticlesRedis(param, results)
		if err != nil {
			log.Printf("upsert to redis error : %s", err.Error())
		}

		return results, nil
	}

	return results, nil
}

func (a *articles) getArticlesDB(param models.ArticleParam) ([]models.Article, error) {
	var (
		results []models.Article
		author  string
		query   string
		where   string
	)
	if param.Author != "" || param.Query != "" {
		where = "WHERE"
	}

	if param.Author != "" {
		author = fmt.Sprintf("author = '%s'", param.Author)
	}

	if param.Query != "" && param.Author != "" {
		pQuery := fmt.Sprintf("%" + param.Query + " %")
		query = fmt.Sprintf("AND body LIKE '%s' AND title LIKE '%s'", pQuery, pQuery)
	} else if param.Query != "" && param.Author == "" {
		pQuery := fmt.Sprintf("%" + param.Query + " %")
		query = fmt.Sprintf("body LIKE '%s' AND title LIKE '%s'", pQuery, pQuery)
	}

	queries := fmt.Sprintf("SELECT * FROM articles %s %s %s", where, author, query)
	log.Print(queries)

	rows, err := a.db.Raw(queries).Rows()
	if err != nil {
		return results, fmt.Errorf("get article error : %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var result models.Article
		if err := rows.Scan(&result.ID, &result.Author, &result.Title, &result.Body, &result.Created); err != nil {
			return results, fmt.Errorf("get article error scan : %s", err.Error())
		}
		results = append(results, result)
		// do something
	}

	return results, nil
}

func (a *articles) getArticlesRedis(param models.ArticleParam) ([]models.Article, error) {
	var (
		results []models.Article
	)

	// serialize query param to string
	rawKey, err := json.Marshal(param)
	if err != nil {
		return results, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}

	// build key
	key := fmt.Sprintf(ArticlesKeyByQuery, string(rawKey))
	// paginationKey := fmt.Sprintf(APIKeyByQuery, string(rawKey))

	resultsRaw, err := a.rd.Get(key).Result()
	if err == redis.Nil {
		return results, err
	} else if err != nil {
		return results, fmt.Errorf("Get cache articles error : %s", err.Error())
	}

	// decode merchant levels (encoded json)
	var decJSON []byte
	decJSON, err = snappy.Decode(decJSON, []byte(resultsRaw))
	if err != nil {
		return results, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}

	// unmarshaling returned byte
	if err := json.Unmarshal(decJSON, &results); err != nil {
		return results, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}
	return results, nil
}

func (a *articles) setArticlesRedis(param models.ArticleParam, v []models.Article) ([]models.Article, error) {

	// serialize query param to string
	rawKey, err := json.Marshal(param)
	if err != nil {
		return v, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}

	// build key
	key := fmt.Sprintf(ArticlesKeyByQuery, string(rawKey))
	// paginationKey := fmt.Sprintf(APIKeyPaginationByQuery, string(rawKey))

	rawJSON, err := json.Marshal(v)
	if err != nil {
		return v, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}

	// snappy compression on merchant statuses
	var encJSON []byte
	encJSON = snappy.Encode(encJSON, rawJSON)

	// set key expiration
	if err := a.rd.Set(key, encJSON, durationArticlesExpiration).Err(); err != nil {
		return v, fmt.Errorf("Marshal cache param error : %s", err.Error())
	}

	return v, nil
}

func (a *articles) createArticleDB(form models.Article) (models.Article, error) {
	row := a.db.Create(&form)

	if row.Error != nil {
		return form, fmt.Errorf("db insert article error : %s", row.Error.Error())
	}

	return form, nil
}
