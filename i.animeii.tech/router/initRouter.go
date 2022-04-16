package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"i.animeii.tech/middleware"
)

func InitRouter(app *gin.Engine) {
	// 静态路由
	app.Use(middleware.Cors())
	app.StaticFS("/", http.Dir("./static/i.animeii.tech"))

}
