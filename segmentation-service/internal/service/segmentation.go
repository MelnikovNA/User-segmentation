package service

import (
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type SegmentationService struct {
	logger     *logrus.Logger
	config     *domain.Config
	repository repository.Repository
}

func NewSegmentationService(p *Params) *SegmentationService {
	return &SegmentationService{
		logger:     p.Logger,
		config:     p.Config,
		repository: *p.Repository}
}
