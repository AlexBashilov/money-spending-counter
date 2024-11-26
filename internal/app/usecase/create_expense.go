package usecase

import (
	"booker/model/apiModels"
	"booker/model/repomodels"
	"context"
	"log"
	"time"
)

func (s *Service) CreateExpense(ctx context.Context, req apiModels.CreateExpenseRequest) error {
	expense := &repomodels.Expense{
		Amount: req.Amount,
		Item:   req.Item,
		Date:   time.Now(),
	}

	_, err := s.itemsRepo.CheckExistItem(ctx, req.Item)
	if err != nil {
		log.Println(err)
	} else {
		if err := s.expenseRepo.CreateExpense(ctx, expense); err != nil {
			return err
		}
	}

	return nil
}
