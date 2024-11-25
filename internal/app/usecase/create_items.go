package usecase

import (
	"booker/internal/app/store"
	"booker/model/apiModels"
	"booker/model/repomodels"
	"context"
)

type Service struct {
	itemsRepo   *store.ItemsRepo
	expenseRepo *store.ExpenseRepo
}

func NewService(itemsRepo *store.ItemsRepo, expenseRepo *store.ExpenseRepo) *Service {
	return &Service{
		itemsRepo:   itemsRepo,
		expenseRepo: expenseRepo,
	}
}

func (s *Service) CreateItems(ctx context.Context, req apiModels.CreateItemsRequest) error {
	U := &repomodels.Items{
		ItemName:    req.ItemName,
		GUID:        req.GUID,
		Description: req.Description,
	}
	if err := s.itemsRepo.CreateItems(ctx, U); err != nil {
		return err
	}
	return nil
}
