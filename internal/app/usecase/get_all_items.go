package usecase

import (
	"booker/model/repomodels"
	"context"
)

func (s *Service) GetAllItems(ctx context.Context) ([]repomodels.Items, error) {
	res, err := s.itemsRepo.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
