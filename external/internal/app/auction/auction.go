package auction

import (
	"context"
	"fmt"
	"net"

	"github.com/gasyok/Assessment/internal/controller/auction"
	"github.com/gasyok/Assessment/internal/repository"
	"github.com/gasyok/Assessment/internal/repository/postgres"
	"github.com/gasyok/Assessment/internal/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/gasyok/Assessment/pkg/auction/v1"
)

type Application struct {
	server *grpc.Server
	lis    net.Listener
}

func New(ctx context.Context, postgresDsn string, interceptors ...grpc.UnaryServerInterceptor) (*Application, error) {
	pool, err := pgxpool.Connect(ctx, postgresDsn)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	txManager := postgres.NewTxManager(pool)
	pgRepo := postgres.NewPgRepository(txManager)
	storage := repository.NewStorage(txManager, pgRepo)
	auctionService := usecase.NewService(storage)
	controller := auction.NewController(auctionService)

	app := &Application{
		server: grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...)),
	}

	reflection.Register(app.server)
	desc.RegisterAuctionServer(app.server, controller)
	return app, nil
}

func (a *Application) WithListener(lis net.Listener) *Application {
	a.lis = lis
	return a
}

func (a *Application) Addr() net.Addr {
	return a.lis.Addr()
}

func (a *Application) Server() *grpc.Server {
	return a.server
}

func (a *Application) Serve() error {
	if err := a.server.Serve(a.lis); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (a *Application) GracefulShutdown() {
	fmt.Println("We trying to gracefully shut down")
	a.server.GracefulStop()
}
