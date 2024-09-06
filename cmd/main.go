package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	api "github.com/xadichamahkamova/fitness-tracking-app/internal/http"
	config "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/load"
	pq "github.com/xadichamahkamova/fitness-tracking-app/internal/pkg/postgres"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/email"
)

func main() {
	
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration loaded")

	db, err := pq.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect to Postgresql")

	notif := email.NewNotificationRepo(*cfg)
	
	queries := storage.New(db)
	r := api.NewGin(queries, *notif)

	addr := fmt.Sprintf(":%s", cfg.ServicePost)
	r.Run(addr)
}
