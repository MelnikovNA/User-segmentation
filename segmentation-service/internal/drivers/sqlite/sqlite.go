package sqlite

import (
	"database/sql"
	"strings"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
)

func New(cfg *domain.Sqlite) (*sql.DB, error) {
	dsn := strings.TrimPrefix(cfg.DSN, "sqlite3://")

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, err
}
