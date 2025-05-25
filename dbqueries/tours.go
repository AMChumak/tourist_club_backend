package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func rows2RouteIds(rows pgx.Rows) ([]model.RouteId, error) {
	var routeIds []model.RouteId
	for rows.Next() {
		routeId := model.RouteId{}
		err := rows.Scan(&routeId.Id)
		if err != nil {
			return nil, fmt.Errorf("convert to routeId model error: %w", err)
		}
		routeIds = append(routeIds, routeId)
	}
	return routeIds, nil
}

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
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByToursCount: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByTour(pg *db.Postgres, ctx context.Context, tour int) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join persons_tours on persons.id = persons_tours.person
			  where tour = @tour`
	args := pgx.NamedArgs{
		"tour": tour,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTour: %w", err)
	}

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
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourTime: %w", err)
	}

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
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourRoute: %w", err)
	}

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
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByTourPlace: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetAllRouteIds(pg *db.Postgres, ctx context.Context) ([]model.RouteId, error) {
	query := `select distinct id
			  from routes`
	rows, err := pg.Db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetToursBySection: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesBySection(pg *db.Postgres, ctx context.Context, section int) ([]model.RouteId, error) {
	query := `select distinct route
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  join persons_roles on persons.id = persons_roles.person
			  where (role = 0 or role = 1) and section = @section`
	args := pgx.NamedArgs{
		"section": section,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetToursBySection: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByTime(pg *db.Postgres, ctx context.Context, fromDate string, toDate string) ([]model.RouteId, error) {
	query := `select distinct route
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  where  (((@toDate::date - tours.start ) >= tours.duration_days) and (((@fromDate::date - tours.start) <= tours.duration_days) or (((@toDate::date - tours.start ) >= 0) and ((@fromDate::date - tours.start) <= 0))))`
	args := pgx.NamedArgs{
		"fromDate": fromDate,
		"toDate":   toDate,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetToursByTime: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByInstructor(pg *db.Postgres, ctx context.Context, instructor int) ([]model.RouteId, error) {
	query := `select distinct route
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  where tours.instructor = @instructor`

	args := pgx.NamedArgs{
		"instructor": instructor,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetToursByInstructor: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByCntGroups(pg *db.Postgres, ctx context.Context, cntGroups int) ([]model.RouteId, error) {
	query := `select distinct id
			  from routes
			  join (
			  select tours.route, count(distinct tours.id) as cnt
			  from persons
			  join persons_tours 
			  on persons.id = persons_tours.person
			  join tours
			  on persons_tours.tour = tours.id
			  group by tours.route) as cnttbl
			  on cnttbl.route = id
			  where cnttbl.cnt >= @cntGroups`

	args := pgx.NamedArgs{
		"cntGroups": cntGroups,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetToursByCntGroups: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByPlace(pg *db.Postgres, ctx context.Context, placeId int) ([]model.RouteId, error) {
	query := `select distinct id 
			  from routes
			  join places_routes 
			  on routes.id = places_routes.route
			  where place = @place`

	args := pgx.NamedArgs{
		"place": placeId,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetRoutesByCntPlace: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByLength(pg *db.Postgres, ctx context.Context, length int) ([]model.RouteId, error) {
	query := `select count( id) 
			  from routes
			  where length_km >= @length`

	args := pgx.NamedArgs{
		"length": length,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetRoutesByLength: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetRoutesByDifficulty(pg *db.Postgres, ctx context.Context, difficulty int) ([]model.RouteId, error) {
	query := `select count( id) 
			  from routes
			  where difficulty >= @difficulty`

	args := pgx.NamedArgs{
		"difficulty": difficulty,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetRoutesByDifficulty: %w", err)
	}

	tours, err := rows2RouteIds(rows)
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func GetAllInstructors(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join tours
			  on tours.instructor = persons.id`

	rows, err := pg.Db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetAllInstructors: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetInstructorsByRole(pg *db.Postgres, ctx context.Context, role int) ([]model.Person, error) {
	query := `select distinct persons.id
			  from persons
			  join tours
			  on tours.instructor = persons.id
			  join persons_roles
			  on persons.id = persons_roles.person
			  where persons_roles.role = @role`

	args := pgx.NamedArgs{
		"role": role,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetInstructorsByRole: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetInstructorsByCategory(pg *db.Postgres, ctx context.Context, routeType int, diffifculty int) ([]model.Person, error) {
	query := `select id, name, surname, patronymic 
			  from( 
			  select persons.id, name, surname, patronymic, max(rt.difficulty) as category
			  from persons
			  join tours
			  on tours.instructor = persons.id
			  join persons_roles
			  on persons.id = persons_roles.person
			  join persons_tours 
			  on persons.id = persons_tours.person join tours as ts on ts.id = persons_tours.tour join routes as rt on rt.id = ts.route
			  where tours.instructor = persons.id  and rt.type= @routeType
			  group by persons.id) 
			  where category >= @difficulty`

	args := pgx.NamedArgs{
		"routeType":  routeType,
		"difficulty": diffifculty,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetInstructorsByCategory: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetInstructorsByCntTours(pg *db.Postgres, ctx context.Context, cntTours int) ([]model.Person, error) {
	query := `select persons.id 
			  from persons 
			  join (
			  select distinct persons.id, count(tours.id) as cnt
			  from persons
			  join tours
			  on tours.instructor = persons.id
			  group by persons.id) as cnttbl 
			  on cnttbl.id = persons.id 
			  where cnt >= @cntTours`

	args := pgx.NamedArgs{
		"cntTours": cntTours,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetInstructorsByCntTours: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetInstructorsByTour(pg *db.Postgres, ctx context.Context, tourId int) ([]model.Person, error) {
	query := `select distinct persons.id                        
			  from persons
			  join tours
			  on tours.instructor = persons.id
			  join routes
			  on routes.id = tours.route
			  where route = @tourId`

	args := pgx.NamedArgs{
		"tourId": tourId,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetInstructorsByTour: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetInstructorsByPlace(pg *db.Postgres, ctx context.Context, placeId int) ([]model.Person, error) {
	query := `select distinct persons.id                        
			  from persons
			  join tours
			  on tours.instructor = persons.id
			  join routes
			  on routes.id = tours.route
			  join places_routes
			  on routes.id = places_routes.route
			  where place = @placeId`

	args := pgx.NamedArgs{
		"placeId": placeId,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetInstructorsByPlaceId: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsWithTrainerInstructor(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select distinct persons.id,name,surname,patronymic
			  from persons 
			  join persons_roles 
			  on persons_roles.person = persons.id 
			  join persons_tours 
			  on persons_tours.person = persons.id
			  join tours
			  on persons_tours.tour = tours.id
			  join groups_persons 
			  on persons.id = groups_persons.person                                                   
			  join groups_workouts 
			  on groups_workouts.group_id = groups_persons.group_id
			  join workout_descriptions
			  on workout_descriptions.id = groups_workouts.workout
			  where workout_descriptions.trainer = tours.instructor`

	rows, err := pg.Db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByToursCount: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsCompletedAll(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select persons.id, name, surname, patronymic 
			  from persons 
			  join (
			  select distinct persons.id, count(distinct route) as cnt                        
			  from persons
			  join persons_roles 
			  on persons_roles.person = persons.id 
			  join persons_tours
			  on persons_tours.person = persons.id 
			  join tours 
			  on tours.id = persons_tours.tour
			  join routes
			  on routes.id = tours.route
			  group by persons.id) as cnttbl 
			  on persons.id = cnttbl.id 
			  where cnttbl.cnt = (select count(*) from routes)`

	rows, err := pg.Db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByToursCount: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsCompletedRoute(pg *db.Postgres, ctx context.Context, routeId int) ([]model.Person, error) {
	query := `select distinct persons.id                        
			  from persons
			  join persons_tours
			  on persons_tours.person = persons.id 
			  join tours 
			  on tours.id = persons_tours.tour
			  join routes
			  on routes.id = tours.route
			  where tours.route = @routeId`

	args := pgx.NamedArgs{
		"routeId": routeId,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByToursCount: %w", err)
	}

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetAllRouteTypes(pg *db.Postgres, ctx context.Context) ([]model.RouteType, error) {
	query := `select id, type from route_types`

	rows, err := pg.Db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetAllRouteTypes: %w", err)
	}
	types, err := rows2RouteType(rows)
	if err != nil {
		return nil, err
	}
	return types, nil
}
