package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"booker/model/apiModels"
	"booker/model/repomodels"
)

func (s *Service) CreateExpense(ctx context.Context, req apiModels.CreateExpenseRequest) error {
	expense := &repomodels.Expense{
		Amount: req.Amount,
		Item:   req.Item,
		Date:   time.Now(),
	}

	existsItemsInDb, err := s.itemsRepo.CheckExistItem(ctx, req.Item)
	if err != nil {
		log.Println(err)
	}
	if !existsItemsInDb {
		return errors.New("статья затрат либо удалена, либо не существует")
	}
	if existsItemsInDb {
		if err := s.expenseRepo.CreateExpense(ctx, expense); err != nil {
			return err
		}
	}

	return nil
}
