package domain

type User struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	ID        uint64  `json:"id"`
	CreatedAt int64   `json:"created_at"`
	Balance   float32 `json:"balance"`
}
