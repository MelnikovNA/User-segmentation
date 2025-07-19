package service

import (
	"context"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type Segmentation interface {
	CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error)
	DeletSegment(ctx context.Context, id int32) (err error)
	UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error)
	GetUserSegments(ctx context.Context, user_id string) ([]*domain.Segmentation, error)
	AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error)
	ListSegments(ctx context.Context, id int32) (listsegments []*domain.Segmentation, err error)
}

type Params struct {
	Logger     *logrus.Logger
	Config     *domain.Config
	Repository *repository.Repository
}

type Services struct {
	Segmentation Segmentation
}

func New(p *Params) *Services {
	return &Services{
		Segmentation: NewSegmentationService(p),
	}
}
