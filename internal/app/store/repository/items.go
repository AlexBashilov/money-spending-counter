package repository

import (
	"booker/model/apiModels"
	"booker/model/repomodels"
	"context"
	"github.com/uptrace/bun"
)

// ItemsRepo initial items repo
type ItemsRepo struct {
	client *bun.DB
}

func NewItemsRepo(client *bun.DB) *ItemsRepo {
	return &ItemsRepo{client: client}
}

func (i *ItemsRepo) GetAllItems() ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemsRepo) GetOnlyOneItem(id int) (*apiModels.UserCostItems, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemsRepo) DeleteItems(id int) error {
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

//
//// GetAllItems get all items
//func (r *ItemsRepo) GetAllItems() ([]map[string]interface{}, error) {
//
//	rows, err := r.store.db.Query(
//		"SELECT id, item_name, guid, description FROM book_cost_items WHERE deleted_at IS NULL",
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	colNames, err := rows.Columns()
//	if err != nil {
//		log.Fatal(err)
//	}
//	cols := make([]interface{}, len(colNames))
//	colPtrs := make([]interface{}, len(colNames))
//	for i := 0; i < len(colNames); i++ {
//		colPtrs[i] = &cols[i]
//	}
//
//	var mySlice = make([]map[string]interface{}, 0)
//	for rows.Next() {
//		var myMap = make(map[string]interface{})
//		err = rows.Scan(colPtrs...)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for i, col := range cols {
//			myMap[colNames[i]] = col
//		}
//		mySlice = append(mySlice, myMap)
//	}
//
//	if len(mySlice) < 1 {
//		return nil, errors.New("no items found")
//	}
//	return mySlice, nil
//}
//
//// DeleteItems delete items
//func (r *ItemsRepo) DeleteItems(id int) error {
//	_, err := r.store.db.Exec("UPDATE public.book_cost_items SET deleted_at = $2 WHERE id = $1;", id, time.Now())
//	if err != nil {
//		log.Fatal(err)
//	}
//	return nil
//}
//
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
