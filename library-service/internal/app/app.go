package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/project/library/config"
	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/controller"
	"github.com/project/library/internal/usecase/library"
	"github.com/project/library/internal/usecase/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repo := repository.NewInMemoryRepository()
	useCases := library.New(logger, repo, repo)

	ctrl := controller.New(logger, useCases, useCases)

	go runRest(ctx, cfg, logger)
	go runGrpc(cfg, logger, ctrl)

	<-ctx.Done()
	const timeToSleeep = time.Second * 3
	time.Sleep(timeToSleeep)
}

func runRest(ctx context.Context, cfg *config.Config, logger *zap.Logger) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	address := "localhost:" + cfg.GRPC.Port
	err := generated.RegisterLibraryHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}

	gatewayPort := ":" + cfg.GatewayPort
	logger.Info("gateway listening at port", zap.String("port", gatewayPort))

	if err = http.ListenAndServe(gatewayPort, mux); err != nil {
		logger.Error("gateway listen error", zap.Error(err))
	}
}

func runGrpc(cfg *config.Config, logger *zap.Logger, libraryService generated.LibraryServer) {
	port := ":" + cfg.GRPC.Port
	lis, err := net.Listen("tcp", port)

	if err != nil {
		logger.Error("can not open tcp socket", zap.Error(err))
		os.Exit(-1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	generated.RegisterLibraryServer(s, libraryService)

	logger.Info("grpc server listening at port", zap.String("port", port))

	if err = s.Serve(lis); err != nil {
		logger.Error("grpc server listen error", zap.Error(err))
	}
}
