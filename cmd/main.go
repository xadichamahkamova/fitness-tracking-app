package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	api "github.com/xadichamahkamova/fitness-tracking-app/internal/http"
	config "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/load"
	pq "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/postgres"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
)

func main() {

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := pq.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	queries := storage.New(db)
	r := api.NewGin(queries)

	addr := fmt.Sprintf(":%s", cfg.Service.Port)
	r.Run(addr)
}
