package sqlite

import (
	"database/sql"
	"strings"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

func New(cfg *domain.Sqlite) (*sql.DB, error) {
	dsn := strings.TrimPrefix(cfg.DSN, "sqlite3://")

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}
