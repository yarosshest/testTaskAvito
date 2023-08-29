package db

import (
	"api/ent"
	"api/ent/segment"
	"api/ent/user"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
)

type QueueUpdateUser struct {
	Add  []string `json:"add"`
	Dell []string `json:"dell"`
}

type UserJson struct {
	Id int `json:"id"`
}

func createUser(client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetID(id).
		Save(ctx)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("user with this id alredy exist: %w", NonUniqueFiledErr)
			}
		}
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func addSegmentToUser(client *ent.Client, u *ent.User, s string) (*ent.User, error) {
	seg, err := client.Segment.Query().Where(
		segment.Name(s),
	).Only(ctx)

	if ent.IsNotFound(err) {
		seg, err = CreateSegment(client, s)
	}

	if err != nil {
		return nil, fmt.Errorf("failed querying segment: %w", err)
	}

	_, err = u.Update().AddSegments(seg).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}
	return u, nil
}

func dellSegmentFromUser(client *ent.Client, u *ent.User, s string) (*ent.User, error) {
	seg, err := client.Segment.Query().Where(
		segment.Name(s),
	).Only(ctx)

	if ent.IsNotFound(err) {
		return nil, fmt.Errorf("failed querying segment: %w", err)
	}

	_, err = u.Update().RemoveSegments(seg).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}
	return u, nil
}

func UpdateUser(client *ent.Client, id int, queue QueueUpdateUser) (*ent.User, error) {
	var u *ent.User
	u, err := client.User.Query().Where(
		user.ID(id),
	).Only(ctx)

	if ent.IsNotFound(err) {
		u, err = createUser(client, id)
	}

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	for _, v := range queue.Add {
		u, err = addSegmentToUser(client, u, v)
		if err != nil {
			return nil, fmt.Errorf("failed updating user: %w", err)
		}
	}

	for _, v := range queue.Dell {
		u, err = dellSegmentFromUser(client, u, v)
		if err != nil {
			return nil, fmt.Errorf("failed updating user: %w", err)
		}
	}

	return u, nil
}

func GetUserSegments(client *ent.Client, id int) ([]string, error) {
	u, err := client.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed fouynd user: %w", NotFoundErr)
	}

	segments, err := u.QuerySegments().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying segments: %w", err)
	}

	segmentsName := make([]string, len(segments))
	for i, v := range segments {
		segmentsName[i] = v.Name
	}

	return segmentsName, nil
}
