package database

import (
	"database/sql"

	"github.com/ias-epf/tutorial/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax) VALUES (?, ?, ?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotalTranscations() (int, error) {
	var total int 
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
