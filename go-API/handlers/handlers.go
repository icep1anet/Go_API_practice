package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go-API/models"
)

// /hello のハンドラ
func HelloHandler(c echo.Context) error {
	name := c.QueryParam("name")
	// gorilla
	fmt.Println(name)
	return c.String(http.StatusOK, "Hello, World")
}

// /article のハンドラ
func PostArticleHandler(c echo.Context) error {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail to encode json\n")
	}
	return c.JSONBlob(http.StatusOK, jsonData)
}

// /article/list のハンドラ
func ArticleListHandler(c echo.Context) error {
	// articles := []models.Article{models.Article1, models.Article2}
	// jsonData, err := json.Marshal(articles)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, "fail to encode json\n")
	// }
	// return c.JSONBlob(http.StatusOK, jsonData)
	query := c.QueryParams()
	var page int = 1
	fmt.Print(query.Get("Page"))
	// if p, ok := query.Get("page"); ok && len(p) > 0 {
	// 	var err error
	// 	page, err = strconv.Atoi(p[0])
	// 	if err != nil {
	// 		return c.String(http.StatusBadRequest, "Invalid path parameter")
	// 	}
	// } else {
	// 	page = 1
	// }
	resString := fmt.Sprintf("Article List (page %d)\n", page)

	return c.String(http.StatusOK, resString)
}

// /article/id のハンドラ
func ArticleDetailHandler(c echo.Context) error {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail to encode json\n")
	}
	return c.JSONBlob(http.StatusOK, jsonData)
}

// /article/nice のハンドラ
func PostNiceHandler(c echo.Context) error {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail to encode json\n")
	}
	return c.JSONBlob(http.StatusOK, jsonData)
}

// /comment のハンドラ
func PostCommentHandler(c echo.Context) error {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail to encode json\n")
	}
	return c.JSONBlob(http.StatusOK, jsonData)
}
