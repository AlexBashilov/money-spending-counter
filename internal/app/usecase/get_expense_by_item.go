package usecase

import (
	"context"

	"booker/model/repomodels"
)

func (s *Service) GetExpenseByItem(ctx context.Context, id int) ([]repomodels.Expense, error) {
	res, err := s.expenseRepo.GetExpenseByItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
