package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func rows2Strain(rows pgx.Rows) ([]model.Strain, error) {
	var strains []model.Strain
	for rows.Next() {
		strain := model.Strain{}
		err := rows.Scan(&strain.Type, &strain.Duration)
		if err != nil {
			return nil, fmt.Errorf("convert to strain model error: %w", err)
		}
		strains = append(strains, strain)
	}
	return strains, nil
}

func GetStrainForTrainer(pg *db.Postgres, ctx context.Context, trainer int, fromDate string, toDate string) ([]model.Strain, error) {
	query := `select  wdat.value as type, sum(finish_time-start_time) as duration 
			  from persons
			  join workout_descriptions as wd
			  on wd.trainer = persons.id
			  join workouts as ws
			  on ws.description = wd.id join workout_descrs_attrs_text as wdat on wdat.descr = wd.id                                  
			  where ws.date between @fromDate and @toDate and persons.id = @trainer
			  group by wdat.value`

	args := pgx.NamedArgs{
		"trainer":  trainer,
		"fromDate": fromDate,
		"toDate":   toDate,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query GetStrainForTrainer: %w", err)
	}

	defer rows.Close()

	strain, err := rows2Strain(rows)
	if err != nil {
		return nil, err
	}

	return strain, nil
}
