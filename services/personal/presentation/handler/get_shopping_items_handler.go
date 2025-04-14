package handler

import (
	"context"
	"log"
	"share-basket/personal/core"
	pb "share-basket/personal/gen/personal"
	"share-basket/personal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GetShoppingItemsHandler interface {
	Handle(ctx context.Context, in *emptypb.Empty) (*pb.GetShoppingItemsResponse, error)
}

type getShoppingItemsHandler struct {
	usecase usecase.GetShoppingItemsUseCase
}

func (g *getShoppingItemsHandler) Handle(ctx context.Context, in *emptypb.Empty) (*pb.GetShoppingItemsResponse, error) {
	outputs, err := g.usecase.Execute(ctx)
	if err != nil {
		return nil, core.NewInvalidError(err)
	}
	log.Println()
	return g.makeResponse(outputs), nil
}

func (g *getShoppingItemsHandler) makeResponse(outputs []usecase.GetShoppingItemOutput) *pb.GetShoppingItemsResponse {
	items := make([]*pb.ShoppingItem, len(outputs))

	for i, v := range outputs {
		var status pb.ShoppingStatus

		switch v.Status {
		case "UnPurchased":
			status = pb.ShoppingStatus_STATUS_UNPURCHASED
		case "Purchased":
			status = pb.ShoppingStatus_STATUS_PURCHASED
		}

		res := &pb.ShoppingItem{
			Id:         v.ID,
			Name:       v.Name,
			Status:     status,
			CategoryId: v.CategoryID,
		}

		items[i] = res
	}

	return &pb.GetShoppingItemsResponse{Items: items}
}

func NewGetShoppingItemsHandler(usecase usecase.GetShoppingItemsUseCase) GetShoppingItemsHandler {
	return &getShoppingItemsHandler{usecase}
}
