package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func rows2Champ(rows pgx.Rows) ([]model.Championship, error) {
	var championships []model.Championship
	for rows.Next() {
		championship := model.Championship{}
		err := rows.Scan(&championship.Id, &championship.Title, &championship.Date)
		if err != nil {
			return nil, fmt.Errorf("convert to championship model error: %w", err)
		}
		championships = append(championships, championship)
	}
	return championships, nil
}

func GetAllChampionships(pg *db.Postgres, ctx context.Context) ([]model.Championship, error) {
	query := `select distinct id, title, date
			  from championships
			  join persons_championships
			  on id = persons_championships.championship
			  join persons_roles 
			  on persons_championships.person = persons_roles.person
			  where extract(day from now() - date) > 0 and role = 1`

	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to do query: %w", err)
	}

	defer rows.Close()

	championships, err := rows2Champ(rows)
	if err != nil {
		return nil, err
	}

	return championships, nil
}

func GetAllChampionshipsBySection(pg *db.Postgres, ctx context.Context, section int) ([]model.Championship, error) {
	query := `select distinct id, title, date
			  from championships
			  join persons_championships
			  on id = persons_championships.championship
			  join persons_roles 
			  on persons_championships.person = persons_roles.person
			  where extract(day from now() - date) > 0 and role = 1 and section = @section`

	args := pgx.NamedArgs{
		"section": section,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query: %w", err)
	}

	defer rows.Close()

	championships, err := rows2Champ(rows)
	if err != nil {
		return nil, err
	}

	return championships, nil
}
