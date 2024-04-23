package api

import (
	"strconv"

	"github.com/KurbanowS/news/internal/app"
	"github.com/KurbanowS/news/internal/models"
	"github.com/gin-gonic/gin"
)

func RatingRoutes(api *gin.RouterGroup) {
	ratingRoutes := api.Group("/ratings")
	{
		ratingRoutes.POST("/likes", LikesCreate)
		ratingRoutes.POST("/dislikes", DislikesCreate)
		ratingRoutes.DELETE("/likes", LikesDelete)
		ratingRoutes.DELETE("/dislikes", DislikesDelete)
	}
}

func LikesCreate(c *gin.Context) {
	r := models.RatingRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	newsIdStr, _ := strconv.Atoi(c.Param("id"))
	newsId := uint(newsIdStr)
	r.NewsId = &newsId
	model, err := app.NewsLikesCreate(&r)
	if err != nil {
		return
	}
	Success(c, gin.H{
		"news_liked": model,
	})
}

func LikesDelete(c *gin.Context) {
	r := models.RatingRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	newsIdStr, _ := strconv.Atoi(c.Param("id"))
	newsId := uint(newsIdStr)
	r.NewsId = &newsId
	model, err := app.NewsLikesDelete(&r)
	if err != nil {
		return
	}
	Success(c, gin.H{
		"like_deleted": model,
	})
}

func DislikesCreate(c *gin.Context) {
	r := models.RatingRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	newsIdStr, _ := strconv.Atoi(c.Param("id"))
	newsId := uint(newsIdStr)
	r.NewsId = &newsId
	model, err := app.NewsDislikesCreate(&r)
	if err != nil {
		return
	}
	Success(c, gin.H{
		"like_deleted": model,
	})
}

func DislikesDelete(c *gin.Context) {
	r := models.RatingRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	newsIdStr, _ := strconv.Atoi(c.Param("id"))
	newsId := uint(newsIdStr)
	r.NewsId = &newsId
	model, err := app.NewsDislikesDelete(&r)
	if err != nil {
		return
	}
	Success(c, gin.H{
		"like_deleted": model,
	})
}
