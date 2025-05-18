package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
	"db_backend/model"
	"strconv"
)

func remove[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func CreatePerson(personReq dto.PersonCreateRequest) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var person model.Person

	person.Name = personReq.Name
	person.Surname = personReq.Surname
	person.Patronymic = personReq.Patronymic

	return dbqueries.InsertPerson(pg, context.Background(), person)
}

func intersection[T comparable](a, b []T) []T {

	for i := 0; i < len(a); {
		resultPerson := a[i]
		found := false
		for _, person := range b {
			if person == resultPerson {
				found = true
				break
			}
		}
		if !found {
			a = remove(a, i)
		} else {
			i++
		}
	}
	return a
}

func checkParameter[T comparable](pg *db.Postgres, parameter string, searchFunc func(pg *db.Postgres, ctx context.Context, section int) ([]T, error), result []T) ([]T, error) {
	if len(result) == 0 {
		return result, nil
	}
	if parameter != "" {

		sectionReady, err := strconv.Atoi(parameter)

		if err != nil {
			return result, err
		}

		resultPart, err := searchFunc(pg, context.Background(), sectionReady)

		if err != nil {
			return result, err
		}
		result = intersection(result, resultPart)
	}
	return result, nil
}

func GetTouristsWithCondition(section string, group string, sex string, birthYear string, age string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllTourists(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetTouristsBySection, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, group, dbqueries.GetTouristsByGroup, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, sex, dbqueries.GetTouristsBySex, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, birthYear, dbqueries.GetTouristsByBirthYear, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, age, dbqueries.GetTouristsByAge, result)
	if err != nil {
		return nil, err
	}

	var response dto.PersonsListResponse

	for _, person := range result {
		var jsonPerson dto.PersonResponse
		jsonPerson.Id = person.Id
		jsonPerson.Name = person.Name
		jsonPerson.Surname = person.Surname
		jsonPerson.Patronymic = person.Patronymic

		response.Persons = append(response.Persons, jsonPerson)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}

func GetTrainersWithCondition(section string, sex string, age string, salary string, specialization string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllTrainers(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetTrainersBySection, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, sex, dbqueries.GetTrainersBySex, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, age, dbqueries.GetTrainersByAge, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, salary, dbqueries.GetTrainersBySalary, result)
	if err != nil {
		return nil, err
	}

	if specialization != "" {
		resultPart, err := dbqueries.GetTrainersBySpecialization(pg, context.Background(), specialization)
		if err != nil {
			return nil, err
		}
		result = intersection(result, resultPart)
	}

	var response dto.PersonsListResponse

	for _, person := range result {
		var jsonPerson dto.PersonResponse
		jsonPerson.Id = person.Id
		jsonPerson.Name = person.Name
		jsonPerson.Surname = person.Surname
		jsonPerson.Patronymic = person.Patronymic

		response.Persons = append(response.Persons, jsonPerson)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}
