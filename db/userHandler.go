package db

import (
	"awesomeProject/ent"
	"awesomeProject/ent/segment"
	"awesomeProject/ent/user"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"go/types"
	"log"
)

type QueueUpdateUser struct {
	Id   int         `json:"id"`
	Add  types.Array `json:"add"`
	Dell types.Array `json:"dell"`
}

func createUser(client *ent.Client, id int) (*ent.User, error) {
	s, err := client.User.
		Create().
		SetName(segment.Name).
		Save(context.Background())

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

func UpdateUser(client *ent.Client, queue QueueUpdateUser) (*ent.User, error) {
	u, err := client.User.Query().Where(
		user.ID(queue.Id),
	).Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("failed querying user: %w", NotFoundErr)
		}
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

}
