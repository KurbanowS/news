package store

import "github.com/KurbanowS/news/internal/models"

type IStore interface {
	NewsFindById(ID string) (*models.News, error)
	NewsFindByIds(Ids []string) ([]*models.News, error)
	NewsFindBy(f models.NewsFilterRequest) (news []*models.News, total int, err error)
	NewsCreate(model *models.News) (*models.News, error)
	NewsUpdate(model *models.News) (*models.News, error)
	NewsDelete(items []*models.News) ([]*models.News, error)

	LikesCreate(newsId uint) error
	LikesDelete(newsId uint) error

	DislikesCreate(newsId uint) error
	DislikesDelete(newsId uint) error
}
