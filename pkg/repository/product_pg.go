package repository

import (
	"context"
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductPostgres struct {
	db *pgxpool.Pool
}

func (p ProductPostgres) GetAllCategories() ([]model.Categories, error) {
	query := "SELECT category FROM item_category"
	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Categories])
	return categories, err
}

func (p ProductPostgres) GetAll() ([]model.Product, error) {

	query := "SELECT item.id,ii.itemname,ii.image,ii.iteminfo,inote.firstprice,ii.garantia,ic.category FROM item" +
		" JOIN item_category ic on ic.id = item.category_id" +
		" JOIN item_info ii on ii.id = item.info_id" +
		" JOIN item_note inote on inote.id = item.note_id"

	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Product])
	return products, err

}

func NewProductPostgres(db *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{db: db}
}
