package usecase

import (
	"context"

	"github.com/halyk/main/internal/domain"
)

type externalAdapter interface {
	Lot(ctx context.Context, sellerID uint64, title string) (*domain.Lot, error)
}
