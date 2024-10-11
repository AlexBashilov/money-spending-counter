package build

import (
	"booker/internal/app/apiserver"
	"booker/internal/app/store/repository"
	"booker/internal/app/usecase"
)

func BuildNewItemsHandler() *apiserver.ItemsHandler {
	bun := NewStore()
	itemsRepo := repository.NewItemsRepo(bun)
	expenseRepo := repository.NewExpenseRepo(bun)
	service := usecase.NewService(itemsRepo, expenseRepo)
	return apiserver.NewItemsHandler(service)
}
