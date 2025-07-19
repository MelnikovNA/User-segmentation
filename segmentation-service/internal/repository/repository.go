package repository

import (
	"context"
	"database/sql"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type Repository interface {
	CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID string, err error)
	DeletSegment(ctx context.Context, id string) (err error)
	UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error)
	GetUserSegments(ctx context.Context, user_id string) ([]string, error)
	AssignRandomSegments(ctx context.Context, id string, percentage float32) (err error)
}

func New(db *sql.DB) Repository {
	return &segmentation{db: db}
}
