package app

import (
	"errors"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/drivers/sqlite"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
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
}
