package dbqueries

import (
	"context"
	"db_backend/db"
	"db_backend/model"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
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

func GetPerson(pg *db.Postgres, ctx context.Context, id int) (*model.Person, error) {
	query := `SELECT * FROM persons WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	var person model.Person
	err := pg.Db.QueryRow(ctx, query, args).Scan(&person.Id, &person.Name, &person.Surname, &person.Patronymic)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func rows2Persons(rows pgx.Rows) ([]model.Person, error) {
	var persons []model.Person
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

func rows2Attributes(rows pgx.Rows) ([]model.Attribute, error) {
	var attrs []model.Attribute
	for rows.Next() {
		attr := model.Attribute{}
		err := rows.Scan(&attr.Id, &attr.Name, &attr.Role, &attr.Type)
		if err != nil {
			return nil, fmt.Errorf("unable to convert row to attribute model: %w", err)
		}
		attrs = append(attrs, attr)
	}
	return attrs, nil
}

func GetPersonRole(pg *db.Postgres, ctx context.Context, person int, section int) (pgtype.Int4, error) {
	query := `SELECT role FROM persons_roles WHERE person = @person AND section = @section`
	args := pgx.NamedArgs{
		"person":  person,
		"section": section,
	}
	row := pg.Db.QueryRow(ctx, query, args)
	var role pgtype.Int4
	err := row.Scan(&role)
	if err != nil {
		return pgtype.Int4{}, err
	}
	return role, nil
}

func DeletePersonRole(pg *db.Postgres, ctx context.Context, person int, section int) error {
	query := `DELETE FROM persons_roles WHERE person = @person AND section = @section`
	args := pgx.NamedArgs{
		"person":  person,
		"section": section,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove old role of user in DeletePersonRole: %w", err)
	}
	return nil
}

func UpdatePersonRole(pg *db.Postgres, ctx context.Context, person int, section int, role int) error {
	query := `DELETE FROM persons_roles WHERE person = @person AND section = @section`
	args := pgx.NamedArgs{
		"person":  person,
		"section": section,
		"role":    role,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove old role of user in UpdatePersonRole: %w", err)
	}

	query = `INSERT INTO persons_roles (person, section, role) VALUES (@person, @section, @role)`

	_, err = pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in UpdatePersonRole: %w", err)
	}
	return nil
}

func GetAllAttributes(pg *db.Postgres, ctx context.Context) ([]model.Attribute, error) {
	query := `SELECT * FROM attributes`
	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve all attributes in GetAllAttributes: %w", err)
	}

	defer rows.Close()

	attrs, err := rows2Attributes(rows)
	if err != nil {
		return nil, err
	}
	return attrs, nil
}

func GetAttribute(pg *db.Postgres, ctx context.Context, id int) (*model.Attribute, error) {
	query := `SELECT * FROM attributes where id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve attribute in GetAttribute: %w", err)
	}

	defer rows.Close()

	attrs, err := rows2Attributes(rows)
	if attrs == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &attrs[0], nil
}

func CreateAttribute(pg *db.Postgres, ctx context.Context, attr model.Attribute) error {
	query := `INSERT INTO attributes (attr, role, attr_type) VALUES (@attr, @role, @attr_type)`
	args := pgx.NamedArgs{
		"attr":      attr.Name,
		"role":      attr.Role,
		"attr_type": attr.Type,
	}
	log.Println(attr.Name, attr.Role, attr.Type)
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in CreateAttribute: %w", err)
	}
	return nil
}

func DeleteAttribute(pg *db.Postgres, ctx context.Context, id int) error {
	query := `DELETE FROM attributes WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove old attribute in DeleteAttribute: %w", err)
	}
	return nil
}

func UpdateAttribute(pg *db.Postgres, ctx context.Context, attr model.Attribute) error {
	query := `UPDATE attributes SET attr = @attr, role = @role, attr_type = @attr_type WHERE id = @id`
	args := pgx.NamedArgs{
		"id":        attr.Id,
		"attr":      attr.Name,
		"role":      attr.Role,
		"attr_type": attr.Type,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update in UpdateAttribute: %w", err)
	}

	return nil
}

func SetPersonIntAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonIntAttribute) error {
	query := `INSERT INTO persons_attrs_int VALUES (@person, @attr, @value)`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in SetPersonIntAttribute: %w", err)
	}
	return nil
}

func SetPersonFloatAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonFloatAttribute) error {
	query := `INSERT INTO persons_attrs_real VALUES (@person, @attr, @value)`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in SetPersonFloatAttribute: %w", err)
	}
	return nil
}

func SetPersonDateAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonDateAttribute) error {
	query := `INSERT INTO persons_attrs_date VALUES (@person, @attr, @value)`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in SetPersonDateAttribute: %w", err)
	}
	return nil
}

func SetPersonStringAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonStringAttribute) error {
	query := `INSERT INTO persons_attrs_text VALUES (@person, @attr, @value)`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row in SetPersonStringAttribute: %w", err)
	}
	return nil
}

func GetPersonIntAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) (int, error) {
	query := `SELECT * FROM persons_attrs_int WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to insert row in GetPersonIntAttribute: %w", err)
	}

	var result int
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}

func GetPersonFloatAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) (float64, error) {
	query := `SELECT * FROM persons_attrs_real WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("unable to insert row in GetPersonFloatAttribute: %w", err)
	}

	var result float64
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}

func GetPersonStringAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) (string, error) {
	query := `SELECT * FROM persons_attrs_text WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return "", fmt.Errorf("unable to insert row in GetPersonStringAttribute: %w", err)
	}

	var result string
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return "", err
		}
	}
	return result, nil
}

func GetPersonDateAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) (pgtype.Date, error) {
	query := `SELECT * FROM persons_attrs_date WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return pgtype.Date{}, fmt.Errorf("unable to insert row in GetPersonDateAttribute: %w", err)
	}

	var result pgtype.Date
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return pgtype.Date{}, err
		}
	}
	return result, nil
}

func UpdatePersonIntAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonIntAttribute) error {
	query := `UPDATE persons_attrs_int SET value = @value WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row in UpdatePersonIntAttribute: %w", err)
	}
	return nil
}

func UpdatePersonFloatAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonFloatAttribute) error {
	query := `UPDATE persons_attrs_real SET value = @value WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row in UpdatePersonFloatAttribute: %w", err)
	}
	return nil
}

func UpdatePersonStringAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonStringAttribute) error {
	query := `UPDATE persons_attrs_text SET value = @value WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row in UpdatePersonStringAttribute: %w", err)
	}
	return nil
}

func UpdatePersonDateAttribute(pg *db.Postgres, ctx context.Context, trio model.PersonDateAttribute) error {
	query := `UPDATE persons_attrs_date SET value = @value WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": trio.PersonId,
		"attr":   trio.AttributeId,
		"value":  trio.Value,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row in UpdatePersonDateAttribute: %w", err)
	}
	return nil
}

func DeletePersonIntAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) error {
	query := `DELETE FROM persons_attrs_int WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row in DeletePersonIntAttribute: %w", err)
	}
	return nil
}

func DeletePersonFloatAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) error {
	query := `DELETE FROM persons_attrs_real WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row in DeletePersonFloatAttribute: %w", err)
	}
	return nil
}

func DeletePersonStringAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) error {
	query := `DELETE FROM persons_attrs_text WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row in DeletePersonStringAttribute: %w", err)
	}
	return nil
}

func DeletePersonDateAttribute(pg *db.Postgres, ctx context.Context, person int, attr int) error {
	query := `DELETE FROM persons_attrs_date WHERE person = @person AND attr = @attr`
	args := pgx.NamedArgs{
		"person": person,
		"attr":   attr,
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row in DeletePersonDateAttribute: %w", err)
	}
	return nil
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
	query := `select distinct id, name, surname, patronymic 
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

func GetTrainersByWorkout(pg *db.Postgres, ctx context.Context, groupNum int, fromDate string, toDate string) ([]model.Person, error) {
	query := `select distinct persons.id, name, surname, patronymic
			  from persons
			  join workout_descriptions as wd
			  on persons.id = wd.trainer
			  join workouts
			  on workouts.description = wd.id
			  join groups_workouts 
			  on groups_workouts.workout = wd.id
			  join groups on groups.id = groups_workouts.group_id
			  where @groupNum = groups.group_number and workouts.date between @from and @to`
	args := pgx.NamedArgs{
		"groupNum": groupNum,
		"from":     fromDate,
		"to":       toDate,
	}
	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to do query GetTrainersByWorkout: %w", err)
	}

	defer rows.Close()

	persons, err := rows2Persons(rows)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetAllManagers(pg *db.Postgres, ctx context.Context) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  where role = 3`

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

func GetManagersBySalary(pg *db.Postgres, ctx context.Context, salary int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_int
			  on persons.id = persons_attrs_int.person
			  where (role = 3) and persons_attrs_int.attr = 6 and persons_attrs_int.value = @salary`
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

func GetManagersByBirthYear(pg *db.Postgres, ctx context.Context, year int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where role = 3 and persons_attrs_date.attr = 2 and @year = extract(year from persons_attrs_date.value)`
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

func GetManagersByAge(pg *db.Postgres, ctx context.Context, age int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where role = 3 and persons_attrs_date.attr = 2 and @age = extract(year from age(persons_attrs_date.value))`
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

func GetManagersByBeginYear(pg *db.Postgres, ctx context.Context, year int) ([]model.Person, error) {
	query := `select distinct id, name, surname, patronymic
			  from persons
			  join persons_roles 
			  on persons.id = persons_roles.person
			  join persons_attrs_date
			  on persons.id = persons_attrs_date.person
			  where role = 3 and persons_attrs_date.attr = 5 and @year = extract(year from persons_attrs_date.value)`
	args := pgx.NamedArgs{
		"year": year,
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
