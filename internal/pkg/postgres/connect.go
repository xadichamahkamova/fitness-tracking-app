package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	config "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/load"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.Postgres)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
