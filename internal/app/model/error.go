package model

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrIncorrectParameters = errors.New("incorrect parameters")

	ErrGroupDuplicate           = errors.New("found group duplicate")
	ErrAllOfficesLinkedToGroup  = errors.New("all offices is linked to group")
	ErrItemsNotValid            = errors.New("items not valid")
	ErrExpenseNotValid          = errors.New("expense not valid")
	ErrItemsAlreadyExist        = errors.New("items already exists")
	ErrItemsNotFound            = errors.New("items not found")
	ErrItemsNotFoundOrExist     = errors.New("items not found or does not exist")
	ErrExpenseRecToDeletedItems = errors.New("can not record expense to deleted items")
)
