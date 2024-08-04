package sqlstore

func (r *BookerRepository) CheckItemIsExist(item string) (bool, error) {
	var exists bool
	err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = $1 AND deleted_at IS NULL)", item).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *BookerRepository) CheckItemIsExistByID(itemID int) (bool, error) {
	var exists bool
	err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE id = $1 AND deleted_at IS NULL)", itemID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

//func (r *BookerRepository) CheckItemIsExistByCode(code int) (bool, error) {
//	var exists bool
//	err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE code = $1 AND deleted_at IS NULL)", code).Scan(&exists)
//	if err != nil {
//		return false, err
//	}
//	return exists, nil
//}

func (r *BookerRepository) CheckExist(comparisonSign interface{}) (bool, error) {
	var exists bool
	switch comparisonSign.(type) {
	case int:
		err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = $1 AND deleted_at IS NULL)", comparisonSign).Scan(&exists)
		if err != nil {
			return false, err
		}
		return exists, nil
	case string:
		var exists bool
		err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = $1 AND deleted_at IS NULL)", comparisonSign).Scan(&exists)
		if err != nil {
			return false, err
		}
		return exists, nil
	default:
		return false, nil
	}
}
