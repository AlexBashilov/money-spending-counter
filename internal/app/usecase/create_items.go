package usecase

import (
	"booker/internal/app/store/repository"
	"booker/model/apiModels"
	"booker/model/repomodels"
)

type Service struct {
	itemsRepo   *repository.ItemsRepo
	expenseRepo *repository.ExpenseRepo
}

func NewService(itemsRepo *repository.ItemsRepo, expenseRepo *repository.ExpenseRepo) *Service {
	return &Service{
		itemsRepo:   itemsRepo,
		expenseRepo: expenseRepo,
	}
}

func (s *Service) CreateItems(req apiModels.CreateItemsRequest) error {
	U := &repomodels.Items{
		ItemName:    req.ItemName,
		GUID:        req.GUID,
		Description: req.Description,
	}
	if err := s.itemsRepo.CreateItems(U); err != nil {
		return err
	}
	return nil
}
