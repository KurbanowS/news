package app

import (
	"errors"
	"strings"

	"github.com/KurbanowS/news/internal/models"
	"github.com/KurbanowS/news/internal/store"
)

func NewsList(f models.NewsFilterRequest) ([]*models.NewsResponse, int, error) {
	news, total, err := store.Store().NewsFindBy(f)
	if err != nil {
		return nil, 0, err
	}
	newsResponse := []*models.NewsResponse{}
	for _, new := range news {
		n := models.NewsResponse{}
		n.FromModel(new)
		newsResponse = append(newsResponse, &n)
	}
	return newsResponse, total, nil
}

func NewsDetail(f models.NewsFilterRequest) (*models.NewsResponse, error) {
	m, _, err := store.Store().NewsFindBy(f)
	if err != nil {
		return nil, err
	}
	if len(m) < 1 {
		return nil, ErrNotFound
	}
	res := &models.NewsResponse{}
	res.FromModel(m[0])
	return res, nil
}

func NewsUpdate(data models.NewsRequest) (*models.NewsResponse, error) {
	model := &models.News{}
	data.ToModel(model)
	var err error
	model, err = store.Store().NewsUpdate(model)
	if err != nil {
		return nil, err
	}
	res := &models.NewsResponse{}
	res.FromModel(model)
	return res, nil
}

func NewsCreate(data models.NewsRequest) (*models.NewsResponse, error) {
	model := &models.News{}
	data.ToModel(model)
	res := &models.NewsResponse{}
	var err error
	model, err = store.Store().NewsCreate(model)
	if err != nil {
		return nil, err
	}
	res.FromModel(model)
	return res, nil
}

func NewsDelete(ids []string) ([]*models.NewsResponse, error) {
	news, err := store.Store().NewsFindByIds(ids)
	if err != nil {
		return nil, err
	}
	if len(news) < 1 {
		return nil, errors.New("model not found: " + strings.Join(ids, ","))
	}
	news, err = store.Store().NewsDelete(news)
	if err != nil {
		return nil, err
	}
	if len(news) == 0 {
		return make([]*models.NewsResponse, 0), nil
	}
	var newsResponse []*models.NewsResponse
	for _, new := range news {
		var n models.NewsResponse
		n.FromModel(new)
		newsResponse = append(newsResponse, &n)
	}
	return newsResponse, nil
}
