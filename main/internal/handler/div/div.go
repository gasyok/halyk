package div

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/halyk/main/internal/domain"
	"github.com/halyk/main/internal/pkg/logger"
	"go.uber.org/zap"
)

type Handler struct {
	divider  divideUsecase
	strategy string
}

func NewHandler(uc divideUsecase, strategy string) *Handler {
	return &Handler{
		divider:  uc,
		strategy: strategy,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Target string `json:"target"`
	}
	flow := func() error {
		req, hash, err := h.ParseRequest(r)
		logger.Info("Hash", zap.Float32("Hash", hash))
		if err != nil {
			return fmt.Errorf("parse: %w", err)
		}

		target, err := h.Division(r.Context(), req, hash, h.strategy)
		if err != nil {
			return fmt.Errorf("h.Division: %w", err)
		}

		resp.Target = target
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			return fmt.Errorf("json.Encode: %w", err)
		}

		return nil
	}

	switch err := flow(); {
	case err == nil:
	case errors.Is(err, domain.ErrInvalidArgument):
		http.Error(w, "invalid argument", http.StatusBadRequest)
	default:
		logger.Error("Internal erro", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Division(ctx context.Context, req *DivideLotRequest, hash float32, strategy string) (string, error) {
	var trg string
	var err error
	switch strategy {
	case "hash":
		trg, err = h.divider.Division(ctx, hash, uint64(req.SellerID), req.Title)
	case "rand":
		trg, err = h.divider.RandDivision(ctx, uint64(req.SellerID), req.Title)
	default:
		return "", fmt.Errorf("strategy error")
	}
	if err != nil {
		return "", fmt.Errorf("divider.Division: %w", err)
	}
	return trg, nil
}
