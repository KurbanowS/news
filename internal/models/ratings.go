package models

type Rating struct {
	ID       uint  `json:"id"`
	NewsId   *uint `json:"news_id"`
	News     *News `json:"news"`
	Likes    *uint `json:"likes"`
	Dislikes *uint `json:"dislikes"`
}

func (Rating) RatingRelationFields() []string {
	return []string{"NewsId"}
}

type RatingResponse struct {
	ID       uint          `json:"id"`
	News     *NewsResponse `json:"news"`
	Likes    *uint         `json:"likes"`
	Dislikes *uint         `json:"dislikes"`
}

type RatingRequest struct {
	ID       *uint `json:"id"`
	NewsId   *uint `json:"news_id"`
	Likes    *uint `json:"likes"`
	Dislikes *uint `json:"dislikes"`
}

func (r *RatingRequest) ToModel(b *Rating) {
	if r.ID != nil {
		b.ID = *r.ID
	}
	b.Likes = r.Likes
	b.Dislikes = r.Dislikes
	if r.NewsId != nil {
		b.NewsId = r.NewsId
	}
}

func (m *RatingResponse) FromModel(b *Rating) {
	m.ID = b.ID
	m.Likes = b.Likes
	m.Dislikes = b.Dislikes
	if m.News != nil {
		m.News = &NewsResponse{}
		m.News.FromModel(b.News)
	}
}

type RatingPaginationRequest struct {
	ID     *uint `form:"id"`
	NewsId *uint `form:"news_id"`
	PaginationRequest
}
