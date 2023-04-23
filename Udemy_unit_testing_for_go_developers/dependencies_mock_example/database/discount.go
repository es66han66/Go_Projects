package database

import "database/sql"

type Repository interface {
	FindCurrentDiscount() (discount int)
}

type DiscountRepository struct {
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (dc *DiscountRepository) FindCurrentDiscount() (discount int) {
	sql := "SELECT value from discount ORDER BY RAND() LIMIT 1"
	row := dc.db.QueryRow(sql)
	row.Scan(&discount)

	return discount
}
