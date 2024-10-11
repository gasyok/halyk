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

	"github.com/halyk/main/internal/adapter/external"
	"github.com/halyk/main/internal/config"
	"github.com/halyk/main/internal/handler"
	"github.com/halyk/main/internal/middleware"
	"github.com/halyk/main/internal/pkg/logger"
	"github.com/halyk/main/internal/usecase"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger.SetLevel(zapcore.InfoLevel)

	baseCfg, err := config.LoadSplit()
	if err != nil {
		logger.Fatal("Failed to load base config", zap.Error(fmt.Errorf("config.LoadSplit: %v", err)))
	}

	externalCfg, err := config.LoadExternal()
	if err != nil {
		logger.Fatal("Failed to load external config", zap.Error(fmt.Errorf("config.LoadExternal: %v", err)))
	}

	externalService := external.NewService(externalCfg.Endpoint(), externalCfg.Timeout)
	copyExternalService := external.NewService(externalCfg.CopyEndpoint(), externalCfg.Timeout)

	usecase := usecase.NewBusinessLogic(baseCfg.Division, externalService, copyExternalService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", baseCfg.Port))
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(fmt.Errorf("net.Listen: %v", err)))
	}

	server := &http.Server{
		Handler: middleware.Chain(
			middleware.WithLogging(),
		)(handler.New(usecase, baseCfg.Strategy)),
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, shutdowner := range []interface{ Shutdown(context.Context) error }{server} {
			if err := shutdowner.Shutdown(shutdownCtx); err != nil {
				logger.Warn("Shutdown error", zap.Error(err))
			}
		}
		logger.Info("Graceful shutdown completed")
	}()

	logger.Info("Starting server", zap.Any("addr", lis.Addr()))
	switch err = server.Serve(lis); {
	case err == nil, errors.Is(err, http.ErrServerClosed):
	default:
		logger.Fatal("Internal server error", zap.Error(err))
	}
}
