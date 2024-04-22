package main

import (
	"log"

	"go-API/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handlers.HelloHandler)
	e.POST("/article", handlers.PostArticleHandler)
	e.GET("/article/list", handlers.ArticleListHandler)
	e.GET("/article/:articleId", handlers.ArticleDetailHandler)
	e.POST("/article/nice", handlers.PostNiceHandler)
	e.POST("/comment", handlers.PostCommentHandler)
	log.Println("server start at port 8080")
	log.Fatal(e.Start(":8080"))
}
