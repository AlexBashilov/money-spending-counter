package usecase

import (
	"context"
)

func (s *Service) DeleteItems(ctx context.Context, id int) error {
	if err := s.itemsRepo.DeleteItems(ctx, id); err != nil {
		return err
	}
	return nil
}
