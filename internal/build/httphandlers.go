package build

import (
	"booker/internal/app/apiserver"
	"booker/internal/app/store"
	"booker/internal/app/usecase"
)

func BuildNewItemsHandler() (*apiserver.ItemsHandler, *apiserver.ExpenseHandler) {
	bun := NewStore()
	itemsRepo := store.NewItemsRepo(bun)
	expenseRepo := store.NewExpenseRepo(bun)
	service := usecase.NewService(itemsRepo, expenseRepo)
	return apiserver.NewItemsHandler(service), apiserver.NewExpenseHandler(service)
}
