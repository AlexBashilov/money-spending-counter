package store

import (
	"booker/model/repomodels"
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"time"
)

// ItemsRepo initial items repo
type ItemsRepo struct {
	client *bun.DB
}

func NewItemsRepo(client *bun.DB) *ItemsRepo {
	return &ItemsRepo{client: client}
}

// CreateItems create item in DB
func (i *ItemsRepo) CreateItems(ctx context.Context, items *repomodels.Items) error {
	_, err := i.client.NewInsert().
		Model(items).
		Column("item_name").
		Column("guid").
		Column("description").
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

// GetAllItems get all items
func (i *ItemsRepo) GetAllItems(ctx context.Context) ([]repomodels.Items, error) {
	var items repomodels.Items
	var result []repomodels.Items
	err := i.client.NewSelect().
		Model(&items).
		Where("deleted_at is null").
		Scan(ctx, &result)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		result = []repomodels.Items{}
	}

	return result, nil
}

// DeleteItems delete items
func (i *ItemsRepo) DeleteItems(ctx context.Context, id int) error {
	var items repomodels.Items

	existsItemsInDb, err := i.CheckExist(ctx, id)
	if err != nil {
		return err
	}
	if existsItemsInDb {
		exist, err := i.CheckItemsDeletedAt(ctx, id)
		if err != nil {
			return err
		}
		if !exist {
			err = i.client.NewUpdate().
				Model(&items).
				Where("id = ?", id).
				Set("deleted_at = ?", time.Now()).
				Scan(ctx)
			if err != nil {
				logrus.Warningln(err)
			}
		} else {
			return errors.New("статья затра удалена и имеет признак deleted_at")
		}
	} else {
		return errors.New("статья затра не существует")
	}

	return nil
}

// GetOne get items by ID
func (i *ItemsRepo) GetOne(ctx context.Context, itemID int) (*repomodels.Items, error) {
	var items repomodels.Items

	err := i.client.NewSelect().
		Model(&items).
		Where("id = ?", itemID).
		Where("deleted_at is null").
		Scan(ctx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Warningln(err)
		return nil, errors.New("статья затрат либо удалена, либо не существует")
	}

	return &items, nil
}

// UpdateItems update items in DB
func (i *ItemsRepo) UpdateItems(ctx context.Context, u *repomodels.Items, id int) error {
	var items repomodels.Items

	existsItemsInDb, err := i.CheckExist(ctx, id)
	if err != nil {
		return err
	}
	if existsItemsInDb {
		exist, err := i.CheckItemsDeletedAt(ctx, id)
		if err != nil {
			return err
		}
		if !exist {
			err = i.client.NewUpdate().
				Model(&items).
				Where("id = ?", id).
				Set("item_name = ?", u.ItemName).
				Set("guid = ?", u.GUID).
				Set("description = ?", u.Description).
				Where("id = ?", u.ID).
				Scan(ctx)
			if err != nil {
				logrus.Warningln(err)
			}
		} else {
			return errors.New("статья затрат удалена и имеет признак deleted_at")
		}
	} else {
		return errors.New("статья затрат не существует")
	}

	return nil
}
