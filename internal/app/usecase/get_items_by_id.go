package usecase

import (
	"booker/model/repomodels"
	"context"
)

func (s *Service) GetItemsByID(ctx context.Context, id int) (*repomodels.Items, error) {
	res, err := s.itemsRepo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
