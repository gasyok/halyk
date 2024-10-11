package usecase

import (
	"context"
	"time"

	"github.com/gasyok/Assessment/internal/domain"
)

type Auction struct {
	repo repository
}

func NewService(repo repository) *Auction {
	return &Auction{
		repo: repo,
	}
}

func (uc *Auction) UpBalance() {
}

func (uc *Auction) CreateLot(
	ctx context.Context,
	sellerID uint64,
	title string,
	description string,
	startPrice float32,
) (*domain.Lot, error) {
	return &domain.Lot{
		SellerID:    sellerID,
		Title:       title,
		Description: description,
		StartPrice:  float64(startPrice),
		CreatedAt:   time.Now(),
	}, nil
}

func (uc *Auction) CreateBid() {
}

func (uc *Auction) CloseAuction() {
}

func (uc *Auction) ProcessTransactions() {
}
