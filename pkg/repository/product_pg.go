package repository

import (
	"course_work/pkg/model"
	"fmt"
	"github.com/jackc/pgx"
)

type ProductPostgres struct {
	db *pgx.Conn
}

func (p ProductPostgres) GetAll() ([]model.Product, error) {
	var products []model.Product
	query := "SELECT item.id, ii.itemname ,ii.itemimage,ii.itemdescription,inote.price,ii.garantia from item" +
		" JOIN item_note inote on inote.id = item.note_id" +
		" JOIN item_info ii on ii.id = item.info_id"

	row := p.db.QueryRow(query)
	fmt.Println(row)
	err := row.Scan(&products)

	return products, err

}

func NewProductPostgres(db *pgx.Conn) *ProductPostgres {
	return &ProductPostgres{db: db}
}
