package repository

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"math/rand/v2"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type segmentation struct {
	db *sql.DB
}

func GetRandomUniqueIDs(totalUsers []int32, percentage float32) []int32 {

	numIDsToReturn := int(math.Floor(float64(float32(len(totalUsers)) * (percentage / 100))))

	rand.Shuffle(len(totalUsers), func(i, j int) {
		totalUsers[i], totalUsers[j] = totalUsers[j], totalUsers[i]
	})

	return totalUsers[:numIDsToReturn]
}

// ListSegments implements Repository.
func (s *segmentation) ListSegments(ctx context.Context, id int32) (listsegments []*domain.Segmentation, err error) {
	var segments []*domain.Segmentation
	if id == -1 {
		stmt := `SELECT id, name FROM segment`
		rows, err := s.db.QueryContext(ctx, stmt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			segment := &domain.Segmentation{}
			err = rows.Scan(&segment.ID, &segment.Name)
			if err != nil {
				return nil, err
			}
			segments = append(segments, segment)
		}
		return segments, err
	} else {
		stmt := `SELECT id, name FROM segment where id=?`
		rows, err := s.db.QueryContext(ctx, stmt, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
		}
		defer rows.Close()

		for rows.Next() {
			segment := &domain.Segmentation{}
			err = rows.Scan(&segment.ID, &segment.Name)
			if err != nil {
				return nil, err
			}
			segments = append(segments, segment)
		}
		return segments, err
	}

}

// AssignRandomSegments implements Repository.
func (s *segmentation) AssignRandomSegments(ctx context.Context, id int32, percentage float32) (err error) {
	stmt := `SELECT id FROM user`
	rows, err := s.db.QueryContext(ctx, stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return err
	}
	defer rows.Close()
	var users []int32

	for rows.Next() {

		var user int32
		err = rows.Scan(&user)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	ids_of_user := GetRandomUniqueIDs(users, percentage)

	ins, err := s.db.PrepareContext(ctx,
		"insert into users_to_segment(user_id, segment_id) values (?,?)")
	if err != nil {
		return err
	}

	for _, user_id := range ids_of_user {

		_, err = ins.ExecContext(ctx, user_id, id)
		if err != nil {
			return err
		}
	}
	return err

}

// CreateSegment implements Repository.
func (s *segmentation) CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error) {
	ins, err := s.db.PrepareContext(ctx,
		"insert into segment (name) values(?)")
	if err != nil {
		return -1, err
	}
	tmp, err := ins.ExecContext(ctx, segmentation.Name)
	if err != nil {
		return -1, err
	}
	res, err := tmp.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int32(res), nil
}

// DeletSegment implements Repository.
func (s *segmentation) DeletSegment(ctx context.Context, id int32) (err error) {
	ins, err := s.db.PrepareContext(ctx, "delete from segment where id=?")
	if err != nil {
		return err
	}

	_, err = ins.ExecContext(ctx, id)
	return err
}

// GetUserSegments implements Repository.
func (s *segmentation) GetUserSegments(ctx context.Context, user_id string) ([]*domain.Segmentation, error) {
	var segments []*domain.Segmentation
	var ids []int32
	stmt := `SELECT segment_id from users_to_segment where user_id=?`
	rows, err := s.db.QueryContext(ctx, stmt, user_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		segment := &domain.Segmentation{}
		err = rows.Scan(&segment.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, segment.ID)
	}

	for _, id := range ids {
		tmp, err := s.ListSegments(ctx, id)
		if err != nil {
			return nil, err
		}
		segments = append(segments, tmp[0])
	}
	return segments, nil
}

// UpdateSegment implements Repository.
func (s *segmentation) UpdateSegment(ctx context.Context, segmentation *domain.Segmentation) (err error) {
	ins, err := s.db.PrepareContext(ctx,
		"update segment set name=? where id=?")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, segmentation.Name, segmentation.ID)
	return err
}
