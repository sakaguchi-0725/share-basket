package handler

import (
	"context"
	"share-basket/personal/core"
	pb "share-basket/personal/proto/gen"
	"share-basket/personal/usecase"
)

type CreateShoppingItemHandler interface {
	Handle(ctx context.Context, req *pb.CreateShoppingItemRequest) (*pb.ShoppingItem, error)
}

type createShoppingItemHandler struct {
	usecase usecase.CreateShoppingItemUseCase
}

func (h *createShoppingItemHandler) Handle(ctx context.Context, req *pb.CreateShoppingItemRequest) (*pb.ShoppingItem, error) {
	input := h.makeInput(req)
	output, err := h.usecase.Execute(ctx, input)
	if err != nil {
		return nil, core.NewInvalidError(err)
	}

	return h.makeResponse(output), nil
}

func (h *createShoppingItemHandler) makeInput(req *pb.CreateShoppingItemRequest) usecase.CreateShoppingItemInput {
	return usecase.CreateShoppingItemInput{
		Name:       req.GetName(),
		CategoryID: req.GetCategoryId(),
	}
}

func (h *createShoppingItemHandler) makeResponse(output usecase.CreateShoppingItemOutput) *pb.ShoppingItem {
	var status pb.ShoppingStatus

	switch output.Status {
	case "UnPurchased":
		status = pb.ShoppingStatus_STATUS_UNPURCHASED
	case "Purchased":
		status = pb.ShoppingStatus_STATUS_PURCHASED
	}

	return &pb.ShoppingItem{
		Id:         output.ID,
		Name:       output.Name,
		Status:     status,
		CategoryId: output.CategoryID,
	}
}

func NewCreateShoppingItemHandler(usecase usecase.CreateShoppingItemUseCase) CreateShoppingItemHandler {
	return &createShoppingItemHandler{usecase}
}
