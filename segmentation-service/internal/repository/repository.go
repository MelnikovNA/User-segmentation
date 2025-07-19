package repository

import (
	"context"
	"database/sql"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type Repository interface {
	CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error)
	DeletSegment(ctx context.Context, id int32) (err error)
	UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error)
	GetUserSegments(ctx context.Context, user_id string) ([]*domain.Segmentation, error)
	AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error)
	ListSegments(ctx context.Context, id int32) (listsegments []*domain.Segmentation, err error)
}

func New(db *sql.DB) Repository {
	return &segmentation{db: db}
}
