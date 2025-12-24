package database

import (
	"context"
	"fmt"
	"log"

	"github.com/ggarnhart/gogogo/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (db *DB) GetRequests(ctx context.Context, page int) ([]models.Request, error) {
	query := `SELECT * FROM REQUESTS LIMIT 20 OFFSET $1`
	rows, err := db.Pool.Query(ctx, query, page)

	if err != nil {
		return nil, fmt.Errorf("error getting requests from DB: %w", err)
	}

	defer rows.Close()

	requests, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Request])

	return requests, nil
}

func (db *DB) UpdateRequestStatus(ctx context.Context, requestId uuid.UUID, updatedStatus string) (int, error) {
	query := `UPDATE REQUESTS SET STATUS = $1, updated_at = NOW() WHERE ID = $2;`
	ctUpdated, err := db.Pool.Exec(ctx, query, updatedStatus, requestId)

	if err != nil {
		return 0, fmt.Errorf("Error updatign requests from DB: %w", err)
	}

	return int(ctUpdated.RowsAffected()), nil
}
