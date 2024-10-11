package domain

type Bid struct {
	ID        uint64  `json:"id"`
	AuctionID uint64  `json:"auction_id"`
	UserID    uint64  `json:"user_id"`
	Amount    float64 `json:"amount"`
	CreatedAt int64   `json:"created_at"`
}
