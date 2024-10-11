package external

type CreateLotRequest struct {
	SellerID   uint64  `json:"seller_id"`
	Title      string  `json:"title"`
	StartPrice float32 `json:"start_price"`
}
