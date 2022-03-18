package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"i.animeii.tech/middleware"
)

func InitRouter(app *gin.Engine) {
	// 静态路由
	app.Use(middleware.Cors())
	app.StaticFile("/", "./static/i.animeii.tech/index.html")
	app.Static("/img", "./static/i.animeii.tech/img")
	app.StaticFS("/images", http.Dir("./static/i.animeii.tech/images"))
	// app.Static("/images", "./static/i.animeii.tech/images")
	app.StaticFile("/favicon.ico", "./static/i.animeii.tech/img/favicon.ico")
	app.Static("/css", "./static/i.animeii.tech/css")
	app.Static("/js", "./static/i.animeii.tech/js")
	app.Static("/post", "./static/i.animeii.tech/post")
	app.Static("/categories", "./static/i.animeii.tech/categories")
	app.Static("/tags", "./static/i.animeii.tech/tags")
	app.Static("/p", "./static/i.animeii.tech/p")
	// anime page
	app.Static("/anime", "./static/i.animeii.tech/anime")
}
