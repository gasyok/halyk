package auction

import (
	"context"

	"github.com/gasyok/Assessment/internal/domain"
)

type auctionService interface {
	UpBalance()
	CreateLot(
		ctx context.Context,
		sellerID uint64,
		title string,
		description string,
		startPrice float32,
	) (*domain.Lot, error)
	CreateBid()
	CloseAuction()
	ProcessTransactions()
}
