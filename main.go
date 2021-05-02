package main

import (
	"article/app/repository"
	"article/app/usecase"
	"article/controller"
	"article/lib/mysql"
	"article/lib/redis"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// Business Layer
	repo *repository.Repository
	uc   *usecase.Usecase

	ctrl controller.Controller
)

func main() {
	// konek to mysql
	db := mysql.GetMysqlConnection()
	rds := redis.Connect()
	// Business layer Initialization
	repo = repository.Init(
		db,
		rds,
	)
	uc = usecase.Init(repo)
	ctrl = controller.Init(uc)
}
