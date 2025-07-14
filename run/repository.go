package run

import (
	"codebox/db"
	"fmt"
	"os"
)

type Repo struct{}

func (r *Repo) CreateResult(entity *Result) error {
	stmt, err := os.ReadFile("./run/create_results_table.sql")

	if err != nil {
		return fmt.Errorf("failed to read create_results_table.sql: %w", err)
	}

	_, err = db.Postgres.Query(string(stmt))
	if err != nil {
		return fmt.Errorf("failed to create results table: %w", err)
	}

	query := `
		INSERT INTO results (request_id, code, language, image, output, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = db.Postgres.Query(
		query,
		entity.RequestId,
		entity.Code,
		entity.Language,
		entity.Image,
		entity.Output,
		entity.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert result: %w", err)
	}

	return nil
}
