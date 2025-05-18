package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func GetTouristsByToursCount(pg *db.Postgres, ctx context.Context, cntTours int) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join (
			  select distinct persons.id, count(tour)
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  where extract(days from (now() - tours.start)) > 0
			  group by persons.id) as cnttbl
			  on persons.id = cnttbl.id
			  where cnttbl.count >=@cntTours`
	args := pgx.NamedArgs{
		"cntTours": cntTours,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByToursCount: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByTour(pg *db.Postgres, ctx context.Context, tour int) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic, group_number
			  from persons
			  join persons_tours on persons.id = persons_tours.person
			  where tour = @tour`
	args := pgx.NamedArgs{
		"tour": tour,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTour: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByTourTime(pg *db.Postgres, ctx context.Context, date string) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join (
			  select distinct persons.id
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  where (@date::date - tours.start) < tours.duration_days and (@date::date - tours.start) >= 0 
			  group by persons.id) as cnttbl
			  on persons.id = cnttbl.id`
	args := pgx.NamedArgs{
		"date": date,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourTime: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByTourRoute(pg *db.Postgres, ctx context.Context, route int) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join (
			  select distinct persons.id
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  where tours.route = @route
			  group by persons.id) as cnttbl
			  on persons.id = cnttbl.id`
	args := pgx.NamedArgs{
		"route": route,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourRoute: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByTourPlace(pg *db.Postgres, ctx context.Context, placeId int) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join (
			  select distinct persons.id
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  join places_routes 
			  on places_routes.route = tours.route
			  where places_routes.place = @place
			  group by persons.id) as cnttbl
			  on persons.id = cnttbl.id`
	args := pgx.NamedArgs{
		"place": placeId,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourPlace: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
