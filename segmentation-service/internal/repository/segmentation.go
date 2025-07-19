package repository

import (
	"context"
	"database/sql"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type segmentation struct {
	db *sql.DB
}

// ListSegments implements Repository.
func (s *segmentation) ListSegments(ctx context.Context, id int32) (listsegments []int32, err error) {
	panic("unimplemented")
}

// AssignRandomSegments implements Repository.
func (s *segmentation) AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error) {
	panic("unimplemented")
}

// CreateSegment implements Repository.
func (s *segmentation) CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error) {
	ins, err := s.db.PrepareContext(ctx,
		"insert into segment")
}

// DeletSegment implements Repository.
func (s *segmentation) DeletSegment(ctx context.Context, id int32) (err error) {
	panic("unimplemented")
}

// GetUserSegments implements Repository.
func (s *segmentation) GetUserSegments(ctx context.Context, user_id string) ([]int32, error) {
	panic("unimplemented")
}

// UpdateSegment implements Repository.
func (s *segmentation) UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error) {
	panic("unimplemented")
}
