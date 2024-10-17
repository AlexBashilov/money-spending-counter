package store

import (
	"booker/model/repomodels"
	"context"
	"fmt"
)

// CheckExist check entity exist in DB
func (i *ItemsRepo) CheckExist(ctx context.Context, id int) (bool, error) {
	var items repomodels.Items

	exists, err := i.client.NewSelect().
		Model(&items).
		Where("id = ?", id).
		Exists(ctx)
	if err != nil {
		return true, fmt.Errorf("ошибка при проверке существования статьи в таблице book_cost_items: %w", err)
	}

	return exists, nil
}

// CheckItemsDeletedAt check entity exist in DB with deleted_at
func (i *ItemsRepo) CheckItemsDeletedAt(ctx context.Context, id int) (bool, error) {
	var items repomodels.Items

	exists, err := i.client.NewSelect().
		Model(&items).
		Where("id = ?", id).
		Where("deleted_at is not null").
		Exists(ctx)
	if err != nil {

		return true, fmt.Errorf("ошибка при проверке существования статьи без даты удаления в таблице book_cost_items: %w", err)
	}

	return exists, nil
}
