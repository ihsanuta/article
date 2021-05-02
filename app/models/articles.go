package models

import "time"

type Article struct {
	ID      int       `json:"id"`
	Author  string    `form:"author" json:"author" binding:"required"`
	Title   string    `form:"title" json:"title" binding:"required"`
	Body    string    `form:"body" json:"body" binding:"required"`
	Created time.Time `json:"created"`
}

type ArticleParam struct {
	Author string `param:"author" db:"author"`
	Query  string `param:"query" db:"query"`
}
