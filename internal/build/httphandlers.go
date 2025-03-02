package build

import (
	"booker/internal/app/apiserver"
	"booker/internal/app/store"
	"booker/internal/app/usecase"
	cfg "booker/internal/config"
)

func BuildNewItemsHandler(cfg cfg.Config) (*apiserver.ItemsHandler, *apiserver.ExpenseHandler) {
	bun := PgsqlConnection(cfg)
	itemsRepo := store.NewItemsRepo(bun)
	expenseRepo := store.NewExpenseRepo(bun)
	service := usecase.NewService(itemsRepo, expenseRepo)
	return apiserver.NewItemsHandler(service), apiserver.NewExpenseHandler(service)
}
