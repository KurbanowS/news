package pgx

import (
	"context"
	"strconv"
	"strings"

	"github.com/KurbanowS/news/internal/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const sqlNewsFields = `n.id, n.article, n.body`
const sqlNewsInsert = `insert into news`
const sqlNewsSelect = `select ` + sqlNewsFields + ` from news n where n.id = ANY($1::int[])`
const sqlNewsSelectMany = `select ` + sqlNewsFields + `, count(*) over() as total from news n where n.id=n.id limit $1 offset $2`
const sqlNewsUpdate = `update news n set id=id`
const sqlNewsDelete = `delete from news n where id = ANY($1::int[])`

func scanNews(rows pgx.Row, m *models.News, addcolumns ...interface{}) (err error) {
	err = rows.Scan(parseColumnForScan(m, addcolumns...)...)
	return
}

func (d *PgxStore) NewsFindById(ID string) (*models.News, error) {
	row, err := d.NewsFindByIds([]string{ID})
	if err != nil {
		return nil, err
	}
	if len(row) < 1 {
		return nil, pgx.ErrNoRows
	}
	return row[0], nil
}

func (d *PgxStore) NewsFindByIds(Ids []string) ([]*models.News, error) {
	news := []*models.News{}
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		rows, err := tx.Query(context.Background(), sqlNewsSelect, (Ids))
		for rows.Next() {
			m := models.News{}
			err := scanNews(rows, &m)
			if err != nil {
				return err
			}
			news = append(news, &m)
		}
		return
	})
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (d *PgxStore) NewsFindBy(f models.NewsFilterRequest) (news []*models.News, total int, err error) {
	args := []interface{}{f.Limit, f.Offset}
	qs, args := NewsListBuildQuery(f, args)
	err = d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		rows, err := tx.Query(context.Background(), qs, args...)
		for rows.Next() {
			new := models.News{}
			err = scanNews(rows, &new, &total)
			if err != nil {
				return err
			}
			news = append(news, &new)
		}
		return
	})
	if err != nil {
		return nil, 0, err
	}
	return news, total, nil
}

func (d *PgxStore) NewsCreate(model *models.News) (*models.News, error) {
	qs, args := NewsCreateQuery(model)
	qs += " RETURNING id"
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		err = tx.QueryRow(context.Background(), qs, args...).Scan(&model.ID)
		return
	})
	if err != nil {
		return nil, err
	}
	editModel, err := d.NewsFindById(strconv.Itoa(int(model.ID)))
	if err != nil {
		return nil, err
	}
	return editModel, nil
}

func (d *PgxStore) NewsUpdate(model *models.News) (*models.News, error) {
	qs, args := NewsUpdateQuery(model)
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), qs, args...)
		return
	})
	if err != nil {
		return nil, err
	}
	editModel, err := d.NewsFindById(strconv.Itoa(int(model.ID)))
	if err != nil {
		return nil, err
	}
	return editModel, nil
}

func (d *PgxStore) NewsDelete(items []*models.News) ([]*models.News, error) {
	ids := []uint{}
	for _, i := range items {
		ids = append(ids, i.ID)
	}
	err := d.runQuery(context.Background(), func(tx *pgxpool.Conn) (err error) {
		_, err = tx.Query(context.Background(), sqlNewsDelete, (ids))
		return
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func NewsCreateQuery(m *models.News) (string, []interface{}) {
	args := []interface{}{}
	cols := ""
	vals := ""
	q := NewsAtomicQuery(m)
	for k, v := range q {
		args = append(args, v)
		cols += ", " + k
		vals += ", $" + strconv.Itoa(len(args))
	}
	qs := sqlNewsInsert + " (" + strings.Trim(cols, ", ") + ") VALUES (" + strings.Trim(vals, ", ") + ")"
	return qs, args
}

func NewsUpdateQuery(m *models.News) (string, []interface{}) {
	args := []interface{}{}
	sets := ""
	q := NewsAtomicQuery(m)
	for k, v := range q {
		args = append(args, v)
		sets += ", " + k + "=$" + strconv.Itoa(len(args))
	}
	args = append(args, m.ID)
	qs := strings.ReplaceAll(sqlNewsUpdate, "set id=id", "set id=id"+sets+" ") + " where id=$" + strconv.Itoa(len(args))
	return qs, args
}

func NewsAtomicQuery(m *models.News) map[string]interface{} {
	q := map[string]interface{}{}
	q["article"] = m.Article
	q["body"] = m.Body
	return q
}

func NewsListBuildQuery(m models.NewsFilterRequest, args []interface{}) (string, []interface{}) {
	var wheres string = ""
	if m.ID != nil && *m.ID != 0 {
		args = append(args, *m.ID)
		wheres += "and n.id=$" + strconv.Itoa(len(args))
	}
	wheres += " order by n.id desc"
	qs := sqlNewsSelectMany
	qs = strings.ReplaceAll(qs, "n.id=n.id", "n.id=n.id"+wheres+" ")
	return qs, args
}
