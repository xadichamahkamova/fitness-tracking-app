package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	config "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/load"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
