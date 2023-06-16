package service

import (
	"WBL0/internal/cache"
	"WBL0/internal/model"
	"WBL0/internal/repository"
)

type Order interface {
	CreateOrder(order model.Order) (string, error)
	GetOrder(uid string) (model.Order, error)
}

type Service struct {
	Order
}

func NewService(repository *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Order: NewOrderService(repository, cache),
	}
}
