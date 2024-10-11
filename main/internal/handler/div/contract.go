package div

import "context"

type divideUsecase interface {
	Division(ctx context.Context, hash float32, sellerID uint64, title string) (string, error)
	RandDivision(ctx context.Context, sellerID uint64, title string) (string, error)
}
