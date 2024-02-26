package router

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", controllers.Index())
	r.GET("/docs", controllers.Docs())
	r.POST("/url/create", controllers.CreateUrl())
	r.GET("/:short_url", controllers.RedirectUrl())

}
