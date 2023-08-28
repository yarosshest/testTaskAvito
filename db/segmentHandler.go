package db

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"testTaskAvito/ent"
	"testTaskAvito/ent/segment"
)

func CreateSegment(client *ent.Client, name string) (*ent.Segment, error) {
	s, err := client.Segment.
		Create().
		SetName(name).
		Save(ctx)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("segment with this name alredy exist: %w", NonUniqueFiledErr)
			}
		}
		return nil, fmt.Errorf("failed creating segment: %w", err)
	}
	log.Println("segment was created: ", s)
	return s, nil

}

func DeleteSegment(client *ent.Client, name string) error {
	_, err := client.Segment.Delete().Where(segment.Name(name)).Exec(ctx)
	switch {
	// If the entity does not meet a specific condition,
	// the operation will return an "ent.NotFoundError".
	case ent.IsNotFound(err):
		return fmt.Errorf("segment was not found: %w", NotFoundErr)
	case err != nil:
		return fmt.Errorf("segment deletion error: %w", err)
	}
	return nil
}
