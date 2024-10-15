package store

import (
	"booker/model/apiModels"
	"booker/model/repomodels"
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"time"
)

// ItemsRepo initial items repo
type ItemsRepo struct {
	client *bun.DB
}

func NewItemsRepo(client *bun.DB) *ItemsRepo {
	return &ItemsRepo{client: client}
}

func (i *ItemsRepo) GetOnlyOneItem(id int) (*apiModels.UserCostItems, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemsRepo) UpdateItems(u *apiModels.UserCostItems, id int) (*apiModels.UserCostItems, error) {
	//TODO implement me
	panic("implement me")
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
	err := i.client.NewUpdate().
		Model(&items).
		Where("id = ?", id).
		Set("deleted_at = ?", time.Now()).
		Scan(ctx)

	if err != nil {
		log.Print(err)
	}

	return nil
}

//// GetOnlyOneItem get items by ID
//func (r *ItemsRepo) GetOnlyOneItem(itemID int) (*model.UserCostItems, error) {
//	var id int
//	var itemName string
//	var guid uuid.UUID
//	var description string
//
//	u := &model.UserCostItems{
//		ID:          id,
//		ItemName:    itemName,
//		GUID:        guid,
//		Description: description,
//	}
//	rows := r.store.db.QueryRow(
//		"SELECT id, item_name, guid, description FROM book_cost_items WHERE id = $1 AND deleted_at IS NULL",
//		itemID,
//	).Scan(
//		&u.ID,
//		&u.ItemName,
//		&u.GUID,
//		&u.Description)
//	if errors.Is(rows, sql.ErrNoRows) {
//		return nil, nil
//	}
//	return u, nil
//}
//
//// UpdateItems update items in DB
//func (r *ItemsRepo) UpdateItems(u *model.UserCostItems, id int) (*model.UserCostItems, error) {
//	_, err := r.store.db.Exec("UPDATE public.book_cost_items SET item_name = $1, guid=$2, description=$3 WHERE id = $4;", u.ItemName, u.GUID, u.Description, id)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return u, nil
//}
