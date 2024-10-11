package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/halyk/main/internal/domain"
)

type Service struct {
	url     *url.URL
	c       *http.Client
	timeout time.Duration
}

func NewService(endpoint string, timeout time.Duration) *Service {
	return &Service{
		url: &url.URL{
			Scheme: "http",
			Host:   endpoint,
		},
		c:       &http.Client{},
		timeout: timeout,
	}
}

func (ps *Service) Lot(ctx context.Context, sellerID uint64, title string) (*domain.Lot, error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(CreateLotRequest{SellerID: uint64(sellerID), Title: title, StartPrice: float32(10)}); err != nil {
		fmt.Println("err in Encoding")
		return nil, domain.ErrServerInternal
	}

	ctx, cancel := context.WithTimeout(ctx, ps.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ps.url.JoinPath("lots").String(), buf)
	if err != nil {
		fmt.Println("err in Request lol")
		return nil, domain.ErrServerInternal
	}
	resp, err := ps.c.Do(req)
	if err != nil {
		fmt.Println("err in Do Request lol")
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, fmt.Errorf("status not found")
	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	lot := new(domain.Lot)
	if err = json.NewDecoder(resp.Body).Decode(lot); err != nil {
		fmt.Println("err in Decoding to Lot")
		return nil, domain.ErrServerInternal
	}
	return lot, nil
}
