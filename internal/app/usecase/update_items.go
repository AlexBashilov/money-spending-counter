package usecase

import (
	"context"

	"booker/model/apiModels"
	"booker/model/repomodels"
)

func (s *Service) UpdateItems(ctx context.Context, req *apiModels.CreateItemsRequest, id int) error {
	U := &repomodels.Items{
		ItemName:    req.ItemName,
		GUID:        req.GUID,
		Description: req.Description,
	}
	if err := s.itemsRepo.UpdateItems(ctx, U, id); err != nil {
		return err
	}
	return nil
}
