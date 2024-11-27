package usecase

import (
	"context"

	"booker/model/repomodels"
)

func (s *Service) GetItemsByID(ctx context.Context, id int) (*repomodels.Items, error) {
	res, err := s.itemsRepo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
