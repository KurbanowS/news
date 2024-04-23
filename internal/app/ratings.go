package app

import (
	"github.com/KurbanowS/news/internal/models"
	"github.com/KurbanowS/news/internal/store"
)

func NewsLikesCreate(model *models.RatingRequest) (*models.RatingResponse, error) {
	var err error
	like := &models.Rating{}
	model.ToModel(like)
	res := &models.RatingResponse{}
	err = store.Store().LikesCreate(*model.NewsId)
	if err != nil {
		return nil, err
	}
	res.FromModel(like)
	return res, nil
}

func NewsDislikesCreate(model *models.RatingRequest) (*models.RatingResponse, error) {
	var err error
	dislike := &models.Rating{}
	model.ToModel(dislike)
	res := &models.RatingResponse{}
	err = store.Store().DislikesCreate(*model.NewsId)
	if err != nil {
		return nil, err
	}
	res.FromModel(dislike)
	return res, nil
}

func NewsLikesDelete(model *models.RatingRequest) (*models.RatingResponse, error) {
	var err error
	like := &models.Rating{}
	model.ToModel(like)
	res := &models.RatingResponse{}
	err = store.Store().DislikesDelete(*model.NewsId)
	if err != nil {
		return nil, err
	}
	res.FromModel(like)
	return res, nil
}

func NewsDislikesDelete(model *models.RatingRequest) (*models.RatingResponse, error) {
	var err error
	dislike := &models.Rating{}
	model.ToModel(dislike)
	res := &models.RatingResponse{}
	err = store.Store().DislikesDelete(*model.NewsId)
	if err != nil {
		return nil, err
	}
	res.FromModel(dislike)
	return res, nil
}
