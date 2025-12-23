package database

import (
	"context"
	"fmt"
	"log"

	"github.com/ggarnhart/gogogo/internal/models"
)

func (db *DB) CreateRequest(ctx context.Context, req *models.CreateRequest) error {
	query := `
		INSERT INTO requests (ip_address, message)
		VALUES ($1, $2);
	`

	_, err := db.Pool.Exec(ctx, query, req.IPAddress, req.Message)

	if err != nil {
		return fmt.Errorf("error creating request in DB: %w", err)
	}

	log.Println("Added request.")
	return nil
}
