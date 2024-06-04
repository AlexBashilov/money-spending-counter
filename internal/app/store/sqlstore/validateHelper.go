package sqlstore

import "fmt"

func (r *BookerRepository) checkItemIsExist(item string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = %s)", item)
	err := r.store.db.QueryRow(query).Scan(&exists)
	return exists, err
}
