package usecase

import "context"

type GetShoppingItemsUseCase interface {
	Execute(ctx context.Context) ([]GetShoppingItemOutput, error)
}

type getShoppingItemsUseCase struct{}

func (g *getShoppingItemsUseCase) Execute(ctx context.Context) ([]GetShoppingItemOutput, error) {
	return []GetShoppingItemOutput{
		{
			ID:         1,
			Name:       "牛乳",
			Status:     "UnPurchased",
			CategoryID: 1,
		},
		{
			ID:         2,
			Name:       "トイレットペーパー",
			Status:     "UnPurchased",
			CategoryID: 2,
		},
	}, nil
}

type GetShoppingItemOutput struct {
	ID         int64
	Name       string
	Status     string
	CategoryID int64
}

func NewGetShoppingItemsUseCase() GetShoppingItemsUseCase {
	return &getShoppingItemsUseCase{}
}
