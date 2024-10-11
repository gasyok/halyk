package handler

import (
	"net/http"

	"github.com/halyk/main/internal/handler/div"
	"github.com/halyk/main/internal/usecase"
)

func New(uc *usecase.BusinessLogic, strategy string) *http.ServeMux {
	divHandler := div.NewHandler(uc, strategy)

	mux := http.NewServeMux()
	mux.Handle("POST /divide", divHandler)

	return mux
}
