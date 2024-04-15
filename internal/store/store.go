package store

import "github.com/KurbanowS/news/internal/store/pgx"

var store IStore

func Store() IStore {
	return store
}

func Init() IStore {
	store = pgx.Init()
	return store
}
