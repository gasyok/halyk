package auction

import (
	"context"
	"errors"
	"fmt"

	"github.com/gasyok/Assessment/internal/controller/auction/serializers"
	"github.com/gasyok/Assessment/internal/domain"
	"github.com/gasyok/Assessment/internal/pkg/logger"
	desc "github.com/gasyok/Assessment/pkg/auction/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) UpBalance(ctx context.Context, req *desc.UpBalanceRequest) (*desc.UpBalanceResponse, error) {
	return new(desc.UpBalanceResponse), nil
}

func (c *Controller) CreateLot(ctx context.Context, req *desc.CreateLotRequest) (*desc.CreateLotResponse, error) {
	var lot *domain.Lot
	flow := func() error {
		var err error
		lot, err = c.auctionService.CreateLot(
			ctx,
			req.SellerId,
			req.Title,
			req.Description,
			req.StartPrice,
		)
		if err != nil {
			return fmt.Errorf("auctionService.CreateLot: %w", err)
		}
		return err
	}
	switch err := flow(); {
	case err == nil:
	case errors.Is(err, domain.ErrInvalidArgument):
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	default:
		logger.Error("Internal error", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return serializers.CreateLotResponse(lot), nil
}
