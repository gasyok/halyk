package usecase

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// Лучше писать что-то смысловое (в нейминге и методах) но для этой задачи можно так
type BusinessLogic struct {
	adapter externalAdapter
	copy    externalAdapter
	div     float32
}

func NewBusinessLogic(div float32, adapter, copy externalAdapter) *BusinessLogic {
	return &BusinessLogic{
		adapter: adapter,
		copy:    copy,
		div:     div,
	}
}

func (uc *BusinessLogic) RandDivision(ctx context.Context, sellerID uint64, title string) (string, error) {
	rand.Seed(uint64(time.Now().UnixNano()))
	randomNumber := rand.Intn(101)
	hash := float32(randomNumber)

	fmt.Println("RANDOM", randomNumber)

	return uc.Division(ctx, hash, sellerID, title)
}

func (uc *BusinessLogic) Division(ctx context.Context, hash float32, sellerID uint64, title string) (string, error) {
	// просто по Хэшу сравниваем
	fmt.Println("WER HERE", hash, uc.div)
	if hash <= uc.div {
		_, err := uc.adapter.Lot(ctx, sellerID, title)
		if err != nil {
			fmt.Println("now error we here")
			return "", fmt.Errorf("adapter.Lot: %w", err)
		}
		fmt.Println("now we here")
		return "api1", nil
	}
	_, err := uc.copy.Lot(ctx, sellerID, title)
	if err != nil {
		return "", fmt.Errorf("copy.Lot: %w", err)
	}
	return "api2", nil
}

//
