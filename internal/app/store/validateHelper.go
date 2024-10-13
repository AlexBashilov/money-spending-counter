package store

// CheckExist check entity exist in DB
//func (r *ItemsRepo) CheckExist(comparisonSign interface{}) (bool, error) {
//	var exists bool
//	switch comparisonSign.(type) {
//	case int:
//		err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE id = $1 AND deleted_at IS NULL)", comparisonSign).Scan(&exists)
//		if err != nil {
//			return false, err
//		}
//		return exists, nil
//	case string:
//		err := r.store.db.QueryRow("SELECT EXISTS(SELECT item_name FROM book_cost_items WHERE item_name = $1 AND deleted_at IS NULL)", comparisonSign).Scan(&exists)
//		if err != nil {
//			return false, err
//		}
//		return exists, nil
//	default:
//		return false, nil
//	}
//}
