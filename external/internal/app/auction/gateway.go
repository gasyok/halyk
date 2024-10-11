package auction

import (
	"context"
	"fmt"
	"net"
	"net/http"

	desc "github.com/gasyok/Assessment/pkg/auction/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	lis    net.Listener
	server *http.Server
}

func (a *Application) Gateway(ctx context.Context, port uint16) (*Gateway, error) {
	conn, err := grpc.NewClient(a.lis.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	mux := runtime.NewServeMux()
	if err = desc.RegisterAuctionHandler(ctx, mux, conn); err != nil {
		return nil, fmt.Errorf("registerLomsHandler: %w", err)
	}

	gateway := new(Gateway)
	gateway.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}
	gateway.server = &http.Server{
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}
	return gateway, nil
}

func (gw *Gateway) Addr() net.Addr {
	return gw.lis.Addr()
}

func (gw *Gateway) Serve() error {
	if err := gw.server.Serve(gw.lis); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (gw *Gateway) Shutdown(ctx context.Context) error {
	if err := gw.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server.Shutdown: %w", err)
	}
	return nil
}
