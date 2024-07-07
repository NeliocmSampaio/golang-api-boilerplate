package router

import (
	"gin-framework-test/basic-api/controllers"
	"gin-framework-test/basic-api/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine

	healthController controllers.HealthController
	bookController   controllers.BookController
}

func NewRouter(
	bookController controllers.BookController,
	healthController controllers.HealthController,
) *Router {
	r := &Router{}
	r.engine = gin.Default()
	r.engine.Use(middlewares.Logger())

	r.bookController = bookController
	r.healthController = healthController

	return r
}

func (r *Router) Run() {
	r.engine.Run()
}
