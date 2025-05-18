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

func rows2Persons(rows pgx.Rows) ([]model.Person, error) {
	persons := []model.Person{}
	for rows.Next() {
		person := model.Person{}
		err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Patronymic)
		if err != nil {
			return nil, fmt.Errorf("unable to convert row to person model: %w", err)
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func GetAllTourists(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic 
              from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  where role = 0 or role = 1`

	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetAllTourists: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsBySection(pg *db.Postgres, ctx context.Context, section int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic               
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  where (role = 0 or role = 1) and section = @section`
	args := pgx.NamedArgs{
		"section": section,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsBySection: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByGroup(pg *db.Postgres, ctx context.Context, groupId int) ([]model.Person, error) {
	query := `select distinct count(*)
			  from persons
			  join groups_persons 
			  on persons.id = groups_persons.person
			  where group_id = @group_id`
	args := pgx.NamedArgs{
		"group_id": groupId,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByGroup: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsBySex(pg *db.Postgres, ctx context.Context, sex int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_int
			  on persons.id = persons_attrs_int.person
			  where (role = 0 or role = 1) and persons_attrs_int.attr = 1 and persons_attrs_int.value = @sex`
	args := pgx.NamedArgs{
		"sex": sex,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsBySex: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByBirthYear(pg *db.Postgres, ctx context.Context, year int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where (role = 0 or role = 1) and persons_attrs_date.attr = 2 and @year = extract(year from persons_attrs_date.value)`
	args := pgx.NamedArgs{
		"year": year,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByBirthYEar: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTouristsByAge(pg *db.Postgres, ctx context.Context, age int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where (role = 0 or role = 1) and persons_attrs_date.attr = 2 and @age = extract(year from age(persons_attrs_date.value))`
	args := pgx.NamedArgs{
		"age": age,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTouristsByAge: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetAllTrainers(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic 
              from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  where role = 2`

	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetAllTrainers: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTrainersBySection(pg *db.Postgres, ctx context.Context, section int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic               
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  where role = 2 and section = @section`
	args := pgx.NamedArgs{
		"section": section,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersBySection: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTrainersBySex(pg *db.Postgres, ctx context.Context, sex int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_int
			  on persons.id = persons_attrs_int.person
			  where role = 2 and persons_attrs_int.attr = 1 and persons_attrs_int.value = @sex`
	args := pgx.NamedArgs{
		"sex": sex,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersBySex: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTrainersByAge(pg *db.Postgres, ctx context.Context, age int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where role = 2 and persons_attrs_date.attr = 2 and @age = extract(year from age(persons_attrs_date.value))`
	args := pgx.NamedArgs{
		"age": age,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersByAge: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTrainersBySalary(pg *db.Postgres, ctx context.Context, salary int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_int
			  on persons.id = persons_attrs_int.person
			  where (role = 2) and persons_attrs_int.attr = 3 and persons_attrs_int.value = @salary`
	args := pgx.NamedArgs{
		"salary": salary,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersBySalary: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetTrainersBySpecialization(pg *db.Postgres, ctx context.Context, specialization string) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_text
		 	  on persons.id = persons_attrs_text.person
		      where (role = 2) and persons_attrs_text.attr = 4 and persons_attrs_text.value = @specialization`
	args := pgx.NamedArgs{
		"specialization": specialization,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersBySpecialization: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
