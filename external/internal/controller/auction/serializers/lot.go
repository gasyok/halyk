package serializers

import (
	"github.com/gasyok/Assessment/internal/domain"
	desc "github.com/gasyok/Assessment/pkg/auction/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateLotResponse(lot *domain.Lot) *desc.CreateLotResponse {
	return &desc.CreateLotResponse{
		Data: &desc.Lot{
			Id:          lot.ID,
			SellerId:    lot.SellerID,
			Title:       lot.Title,
			Description: lot.Description,
			StartPrice:  float32(lot.StartPrice),
			AuctionId:   lot.AuctionID,
			CreatedAt:   timestamppb.New(lot.CreatedAt),
		},
	}
}
