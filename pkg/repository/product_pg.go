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

func (p ProductPostgres) GetAllCategories() ([]model.ProductCategory, error) {
	query := "SELECT _category FROM item_category"
	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.ProductCategory])

	return categories, err
}

func (p ProductPostgres) GetAll() ([]model.Product, error) {

	query := "SELECT item.id, ii.itemname ,ii.itemimage,ii.itemdescription,inote.price,ii.garantia,ic._category from item" +
		" JOIN item_note inote on inote.id = item.note_id" +
		" JOIN item_info ii on ii.id = item.info_id" +
		" JOIN item_category ic on ic.id = item.category_id"

	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	items, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Product])

	return items, err

}

func NewProductPostgres(db *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{db: db}
}
