package repository

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

type segmentation struct {
	db *sql.DB
}

func GetRandomUniqueIDs(totalUsers int, percentage float32) []int {
	numIDsToReturn := int(math.Floor(float64(581 * (percentage / 100))))

	if totalUsers <= 0 {
		return []int{}
	}

	rand.Seed(time.Now().UnixNano())

	allIDs := make([]int, totalUsers)
	for i := 0; i < totalUsers; i++ {
		allIDs[i] = i + 1
	}

	for i := len(allIDs) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		allIDs[i], allIDs[j] = allIDs[j], allIDs[i]
	}

	return allIDs[:numIDsToReturn]
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
		rows, err := s.db.QueryContext(ctx, stmt)
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
	ids_of_user := GetRandomUniqueIDs(581, percentage)
	for _, id := range ids_of_user {
		ins, err := s.db.PrepareContext(ctx,
		"insert into users_to_segment(user_id,email,password) values (?,?,?)"))
	}

}

// CreateSegment implements Repository.
func (s *segmentation) CreateSegment(ctx context.Context, segmentation *domain.Segmentation) (segmentID int32, err error) {
	ins, err := s.db.PrepareContext(ctx,
		"insert into segment (name) values(?)")
	if err != nil {
		return -1, err
	}
	tmp, err := ins.ExecContext(ctx, segmentation.ID)
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
	ins, err := s.db.PrepareContext(ctx, "delete from user where id=?")
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
	stmt := `SELECT id, deck_name,  from deck_descr where user_id=?`
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
