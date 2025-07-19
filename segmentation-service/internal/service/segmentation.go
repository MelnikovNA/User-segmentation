package service

import (
	"context"

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

func (s *SegmentationService) CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error) {
	panic("unimplemented")
}

func (s *SegmentationService) AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error) {
	panic("unimplemented")
}

func (s *SegmentationService) DeletSegment(ctx context.Context, id int32) (err error) {
	panic("unimplemented")
}

func (s *SegmentationService) GetUserSegments(ctx context.Context, user_id string) ([]string, error) {
	panic("unimplemented")
}

func (s *SegmentationService) UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error) {
	panic("unimplemented")
}

func (s *SegmentationService) ListSegments(ctx context.Context, id int32) (listsegments []int32, err error) {
	panic("unimplemented")
}
