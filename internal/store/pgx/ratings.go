package pgx

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const sqlLikesFields = `insert into ratings(l.news_id, l.likes) values($1, $2)`
const sqlDislikesFields = `insert into ratings(d.news_id, l.dislikes) values($1, $2)`
const sqlNewsRatingDelete = `delete from rating r where r.news_id = $1`

func (d *PgxStore) LikesCreate(newsId uint) error {
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), sqlLikesFields, newsId)
		return
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *PgxStore) DislikesCreate(newsId uint) error {
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), sqlDislikesFields, newsId)
		return
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *PgxStore) LikesDelete(newsId uint) error {
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), sqlNewsRatingDelete, newsId)
		return
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *PgxStore) DislikesDelete(newsId uint) error {
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), sqlNewsRatingDelete, newsId)
		return
	})
	if err != nil {
		return err
	}
	return nil
}
