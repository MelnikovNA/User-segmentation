package repository

import (
	"context"
	"database/sql"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type segmentation struct {
	db *sql.DB
}

// AssignRandomSegments implements Repository.
func (s *segmentation) AssignRandomSegments(ctx context.Context, id string, percentage float32) (err error) {
	panic("unimplemented")
}

// CreateSegment implements Repository.
func (s *segmentation) CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID string, err error) {
	panic("unimplemented")
}

// DeletSegment implements Repository.
func (s *segmentation) DeletSegment(ctx context.Context, id string) (err error) {
	panic("unimplemented")
}

// GetUserSegments implements Repository.
func (s *segmentation) GetUserSegments(ctx context.Context, user_id string) ([]string, error) {
	panic("unimplemented")
}

// UpdateSegment implements Repository.
func (s *segmentation) UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error) {
	panic("unimplemented")
}
