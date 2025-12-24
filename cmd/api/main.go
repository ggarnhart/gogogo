package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ggarnhart/gogogo/internal/config"
	"github.com/ggarnhart/gogogo/internal/database"
	"github.com/ggarnhart/gogogo/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.New(ctx)

	if err != nil {
		log.Fatal("Could not connect to db", err)
	}
	defer db.Close(ctx)

	requestHandler := handlers.NewRequestHandler(db)

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handlers.HealthHandler)
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/request", requestHandler.CreateRequestHandler)
		r.Get("/requests", requestHandler.GetRequestsHandler)
	})

	log.Println("current port", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
