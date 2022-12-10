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

func (p ProductPostgres) PostProductToStock(product model.ProductToAdd) error {
	query := "INSERT INTO item_note(date_of_delivery, firstPrice, firstTax)" +
		"VALUES (now(),$1,15) RETURNING id;"
	var itemNoteId int
	err := p.db.QueryRow(context.TODO(), query, product.Price).Scan(&itemNoteId)
	if err != nil {
		return err
	}
	query = "INSERT INTO item_info(itemname, iteminfo, image, garantia)" +
		"VALUES ($1,$2,$3,$4) RETURNING id;"
	var itemInfoId int
	err = p.db.QueryRow(context.TODO(), query, product.Name, product.Description, product.Image, product.Garantia).Scan(&itemInfoId)
	if err != nil {
		return err
	}
	query = "INSERT INTO item(note_id, info_id, provider_id, category_id)" +
		"VALUES ($1,$2,$3,$4) RETURNING id;"
	var itemId int
	err = p.db.QueryRow(context.TODO(), query, itemNoteId, itemInfoId, product.ProviderId, product.CategoryId).Scan(&itemId)
	if err != nil {
		return err
	}
	query = "INSERT INTO mainstock(item_id, itemcount)" +
		"VALUES($1,$2)"
	_, err = p.db.Exec(context.TODO(), query, itemId, product.Count)
	if err != nil {
		return err
	}
	return err
}

func (p ProductPostgres) DeleteById(id int) error {
	query := "DELETE FROM item WHERE id=$1"
	_, err := p.db.Query(context.TODO(), query, id)

	return err
}

func (p ProductPostgres) GetAllProviders() ([]model.Providers, error) {
	query := "SELECT id,name FROM provider"
	rows, err := p.db.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Providers])
	return categories, err
}

func (p ProductPostgres) GetAllCategories() ([]model.Categories, error) {
	query := "SELECT * FROM item_category"
	rows, err := p.db.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Categories])
	return categories, err
}

func (p ProductPostgres) GetAll() ([]model.Stock, error) {

	query := "\nSELECT itemCount,i.id,ii.itemname,ii.image,ii.iteminfo,inote.firstPrice,ii.garantia,ic.category,p.name FROM mainStock\nJOIN item i on i.id = mainStock.item_id\nJOIN item_info ii on ii.id = i.info_id\nJOIN item_category ic on ic.id = i.category_id\nJOIN item_note inote on inote.id = i.note_id\nJOIN provider p on p.id = i.provider_id\n"

	rows, err := p.db.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Stock])
	return products, err
}

func NewProductPostgres(db *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{db: db}
}
