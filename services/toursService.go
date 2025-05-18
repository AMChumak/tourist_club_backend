package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
)

func GetTouristsByTour(section string, group string, cntTours string, tourId string, tourTime string, routeId string, placeId string) (*dto.PersonsListResponse, error) {
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
	result, err = checkParameter(pg, cntTours, dbqueries.GetTouristsByToursCount, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, tourId, dbqueries.GetTouristsByTour, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, routeId, dbqueries.GetTouristsByTourRoute, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, placeId, dbqueries.GetTouristsByTourPlace, result)
	if err != nil {
		return nil, err
	}

	if tourTime != "" {
		resultPart, err := dbqueries.GetTouristsByTourTime(pg, context.Background(), tourTime)
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
