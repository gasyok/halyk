package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gasyok/Assessment/internal/app/auction"
	"github.com/gasyok/Assessment/internal/config"
	"github.com/gasyok/Assessment/internal/interceptor"
	"github.com/gasyok/Assessment/internal/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	logger.SetLevel(zapcore.InfoLevel)

	auctionCfg, err := config.LoadAuction()
	if err != nil {
		logger.Fatal("Failed to load auction config", zap.Error(fmt.Errorf("config.LoadAuction: %v", err)))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", auctionCfg.GRPCPort))
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(fmt.Errorf("net.Listen: %v", err)))
	}

	app, err := auction.New(
		ctx,
		auctionCfg.ConnString(),
		interceptor.WithValidation(),
		interceptor.WithLogging(),
	)
	if err != nil {
		logger.Fatal("Failed to initialize app", zap.Error(err))
	}

	app.WithListener(lis)

	gw, err := app.Gateway(ctx, auctionCfg.GatewayPort)
	if err != nil {
		logger.Fatal("Failed to initialize gateway", zap.Error(err))
	}

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, shutdowner := range []interface{ Shutdown(context.Context) error }{gw} {
			if err := shutdowner.Shutdown(shutdownCtx); err != nil {
				logger.Warn("Shutdown error", zap.Error(err))
			}
		}
		app.GracefulShutdown()
		logger.Info("Graceful shutdown completed")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Info("Starting HTTP gateway server", zap.Any("addr", gw.Addr()))
		switch err := gw.Serve(); {
		case err == nil, errors.Is(err, http.ErrServerClosed):
		default:
			logger.Error("Internal HTTP gateway server error", zap.Error(err))
		}
	}()

	logger.Info("Starting gRPC server", zap.Any("addr", app.Addr()))
	if err := app.Serve(); err != nil {
		logger.Fatal("Internal gRPC server error", zap.Error(err))
	}
}
