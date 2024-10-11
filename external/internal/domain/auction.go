package domain

type Auction struct {
	ID        uint64 `json:"id"`
	LotID     uint64 `json:"lot_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	WinnerID  uint64 `json:"winner_id"`
	CreatedAt int64  `json:"created_at"`
}
