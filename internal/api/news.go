package api

import (
	"strconv"

	"github.com/KurbanowS/news/internal/app"
	"github.com/KurbanowS/news/internal/models"
	"github.com/gin-gonic/gin"
)

func NewsRoutes(api *gin.RouterGroup) {
	newsRoutes := api.Group("/news")
	{
		newsRoutes.GET("", NewsList)
		newsRoutes.GET(":id", NewsDetail)
		newsRoutes.POST("", NewsCreate)
		newsRoutes.PUT(":id", NewsUpdate)
		newsRoutes.DELETE("", NewsDelete)
	}
}

func NewsList(c *gin.Context) {
	r := models.NewsFilterRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	news, total, err := app.NewsList(r)
	if err != nil {
		handleError(c, err)
		return
	}
	Success(c, gin.H{
		"news":  news,
		"total": total,
	})
}

func NewsDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	idu := uint(id)
	if id == 0 {
		handleError(c, err)
		return
	}

	args := models.NewsFilterRequest{
		ID: &idu,
	}
	res, err := app.NewsDetail(args)
	if err != nil {
		handleError(c, err)
		return
	}
	Success(c, gin.H{
		"news_detail": res,
	})
}

func NewsUpdate(c *gin.Context) {
	r := models.NewsRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	idp := uint(id)
	r.ID = &idp

	if id == 0 {
		handleError(c, app.ErrRequired.SetKey("id"))
		return
	}
	news, err := app.NewsUpdate(r)
	if err != nil {
		handleError(c, err)
		return
	}
	Success(c, gin.H{
		"updated_news": news,
	})
}

func NewsCreate(c *gin.Context) {
	r := models.NewsRequest{}
	if errMsg, errKey := BindAndValidate(c, &r); errMsg != "" || errKey != "" {
		handleError(c, app.NewAppError(errMsg, errKey, ""))
		return
	}
	news, err := app.NewsCreate(r)
	if err != nil {
		handleError(c, err)
		return
	}
	Success(c, gin.H{
		"created_news": news,
	})
}

func NewsDelete(c *gin.Context) {
	var ids []string = c.QueryArray("ids")
	if len(ids) == 0 {
		handleError(c, app.ErrRequired.SetKey("ids"))
		return
	}
	news, err := app.NewsDelete(ids)
	if err != nil {
		handleError(c, err)
		return
	}
	Success(c, gin.H{
		"deleted_news": news,
	})
}
