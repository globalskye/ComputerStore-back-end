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

func (p ProductPostgres) GetAllProviders() ([]model.Providers, error) {
	query := "SELECT name FROM provider"
	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Providers])
	return categories, err
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

func (p ProductPostgres) GetAll() ([]model.Stock, error) {

	query := "\nSELECT itemCount,i.id,ii.itemname,ii.image,ii.iteminfo,inote.firstPrice,ii.garantia,ic.category,p.name FROM mainStock\nJOIN item i on i.id = mainStock.item_id\nJOIN item_info ii on ii.id = i.info_id\nJOIN item_category ic on ic.id = i.category_id\nJOIN item_note inote on inote.id = i.note_id\nJOIN provider p on p.id = i.provider_id\n"

	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Stock])
	return products, err
}

func NewProductPostgres(db *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{db: db}
}
