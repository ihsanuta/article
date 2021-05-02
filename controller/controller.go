package controller

import (
	usecase "article/app/usecase"
	"article/config"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

type Controller interface{}

var once = &sync.Once{}

type controller struct {
	usecase *usecase.Usecase
}

func Init(usecase *usecase.Usecase) Controller {
	var ctrl *controller
	once.Do(func() {
		ctrl = &controller{
			usecase: usecase,
		}
		ctrl.Serve()
	})
	return ctrl
}

func (c *controller) Serve() {
	router := gin.Default()
	group := router.Group("/api/v1")
	group.GET("/articles", c.GetArticles)
	group.POST("/articles", c.CreateArticles)

	serverString := fmt.Sprintf("%s:%s", config.AppConfig["host"], config.AppConfig["port"])
	router.Run(serverString)

}
