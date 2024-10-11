package domain

import "time"

type Lot struct {
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartPrice  float64   `json:"start_price"`
	ID          uint64    `json:"id"`
	SellerID    uint64    `json:"seller_id"`
	AuctionID   uint64    `json:"auction_id"`
}
