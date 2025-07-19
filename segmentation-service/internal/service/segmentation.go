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
	seg_id, err := s.repository.CreateSegment(ctx, segmentation)
	if err != nil {
		return -1, err
	}
	return seg_id, err

}

func (s *SegmentationService) AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error) {
	err = s.repository.AssignRandomSegments(ctx, id, percentage)
	if err != nil {
		return err
	}
	return nil
}

func (s *SegmentationService) DeletSegment(ctx context.Context, id int32) (err error) {
	err = s.repository.DeletSegment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SegmentationService) GetUserSegments(ctx context.Context, user_id string) ([]string, error) {
	segments, err := s.repository.GetUserSegments(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return segments, err
}

func (s *SegmentationService) UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error) {
	err = s.repository.UpdateSegment(ctx, segmentation)
	if err != nil {
		return err
	}
	return nil
}

func (s *SegmentationService) ListSegments(ctx context.Context, id int32) (listsegments []int32, err error) {
	listsegments, err = s.repository.ListSegments(ctx, id)
	if err != nil {
		return nil, err
	}
	return listsegments, nil
}
