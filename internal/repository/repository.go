package repository

import (
	"WBL0/internal/model"
	"WBL0/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order model.Order) (string, error)
	GetOrder(uid string) (model.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: postgres.NewOrderPostgres(db),
	}
}
