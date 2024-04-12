package service

import (
	"context"
	"effective_mobile_task/internal/model"
	"effective_mobile_task/internal/repository"
)

type Car interface {
	Create(ctx context.Context, items []model.Car) error
	Update(ctx context.Context, itemID uint, item *model.CarUpdate) error
	Delete(ctx context.Context, itemID uint) error
	Get(ctx context.Context, limit, offset int) ([]model.Car, error)
}

type car struct {
	repo repository.Car
}

func CarService(repo repository.Car) Car {
	return &car{repo: repo}
}

func (s *car) Create(ctx context.Context, items []model.Car) error {
	return s.repo.Create(ctx, items)
}

func (s *car) Update(ctx context.Context, itemID uint, item *model.CarUpdate) error {
	return s.repo.Update(ctx, itemID, item)
}

func (s *car) Delete(ctx context.Context, itemID uint) error {
	return s.repo.Delete(ctx, itemID)
}

func (s *car) Get(ctx context.Context, limit, offset int) ([]model.Car, error) {
	return s.repo.Get(ctx, limit, offset)
}
