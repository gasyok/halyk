package div

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"math/big"
	"net/http"

	"github.com/halyk/main/internal/domain"
)

type InspectCartRequest struct {
	UserID uint64 `json:"-"`
}

type DivideLotRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartPrice  float64 `json:"start_price"`
	SellerID    uint64  `json:"seller_id"`
	AuctionID   uint64  `json:"auction_id"`
}

func (*Handler) ParseRequest(r *http.Request) (*DivideLotRequest, float32, error) {
	var (
		request = new(DivideLotRequest)
		err     error
	)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, 0, domain.ErrServerInternal
	}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, 0, domain.ErrInvalidArgument
	}
	floatVal, err := calculateFloatHash(r, body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to calculate request hash: %v", err)
	}

	return request, floatVal, nil
}

func calculateFloatHash(r *http.Request, body []byte) (float32, error) {
	h := sha256.New()

	addHeaderToHash(h, r, "Content-Type")
	addHeaderToHash(h, r, "Authorization")

	h.Write(body)

	hashValue := hex.EncodeToString(h.Sum(nil))

	intHash := new(big.Int)
	intHash.SetString(hashValue[:16], 16)

	mod := new(big.Int)
	mod.Mod(intHash, big.NewInt(100))

	floatVal := float32(mod.Int64()) + 1

	return floatVal, nil
}

func addHeaderToHash(h hash.Hash, r *http.Request, headerName string) {
	if value := r.Header.Get(headerName); value != "" {
		h.Write([]byte(value))
	}
}
