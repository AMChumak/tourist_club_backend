package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func InsertPerson(pg *db.Postgres, ctx context.Context, person model.Person) error {
	query := `INSERT INTO persons (name,surname,patronymic) VALUES (@name, @surname, @patronymic)`
	args := pgx.NamedArgs{
		"name":       person.Name,
		"surname":    person.Surname,
		"patronymic": person.Patronymic,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
