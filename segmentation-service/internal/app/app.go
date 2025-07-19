package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/drivers/sqlite"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/repository"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	ErrSignalReceived = errors.New("signal received")
)

func Run(config string) error {
	cfg := &domain.Config{}

	if err := cleanenv.ReadConfig(config, cfg); err != nil {
		return err
	}

	parseFlags(cfg)

	log := logrus.New()
	log.Infof("Hello app!%#v", cfg)

	db, err := sqlite.New(&cfg.Sqlite)
	if err != nil {
		return err
	}
	return err

	r := repository.New(db)
	s := service.New(&service.Params{
		Logger:     log,
		Config:     cfg,
		Repository: &r,
	})
	eg, egctx := errgroup.WithContext(context.Background())
	g := grpcserver.New(cfg.GrpcServer.Host, cfg.GrpcServer.Port, s, log)

	_ = egctx
	_ = g

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	eg.Go(func() error {
		log.Infof("Grpc listener has started on %s:%s", cfg.GrpcServer.Host, cfg.GrpcServer.Port)
		return g.Serve()
	})
}
