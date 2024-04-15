package models

type News struct {
	ID      uint    `json:"id"`
	Article *string `json:"article"`
	Body    *string `json:"body"`
}

func (News) RelationFields() []string {
	return []string{}
}

type NewsRequest struct {
	ID      *uint   `json:"id"`
	Article *string `json:"article"`
	Body    *string `json:"body"`
}

type NewsResponse struct {
	ID      uint    `json:"id"`
	Article *string `json:"article"`
	Body    *string `json:"body"`
}

func (b *NewsRequest) ToModel(m *News) {
	if b.ID != nil {
		m.ID = *b.ID
	}
	m.Article = b.Article
	m.Body = b.Body
}

func (r *NewsResponse) FromModel(m *News) {
	r.ID = m.ID
	r.Article = m.Article
	r.Body = m.Body
}

type NewsFilterRequest struct {
	ID *uint `form:"id"`
	PaginationRequest
}
