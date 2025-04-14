package usecase

import "context"

type CreateShoppingItemUseCase interface {
	Execute(ctx context.Context, input CreateShoppingItemInput) (CreateShoppingItemOutput, error)
}

type createShoppingItemUseCase struct{}

func (c *createShoppingItemUseCase) Execute(ctx context.Context, input CreateShoppingItemInput) (CreateShoppingItemOutput, error) {
	return CreateShoppingItemOutput{
		ID:         1,
		Name:       "牛乳",
		Status:     "UnPurchased",
		CategoryID: 1,
	}, nil
}

type CreateShoppingItemInput struct {
	Name       string
	CategoryID int64
}

type CreateShoppingItemOutput struct {
	ID         int64
	Name       string
	Status     string
	CategoryID int64
}

func NewCreateShoppingItemUseCase() CreateShoppingItemUseCase {
	return &createShoppingItemUseCase{}
}
