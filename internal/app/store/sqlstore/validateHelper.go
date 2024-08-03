package sqlstore

func (r *BookerRepository) CheckItemIsExist(item string) (bool, error) {
	var exists bool
	err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = $1)", item).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
