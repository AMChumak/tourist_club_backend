package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
	"db_backend/model"
	"strconv"
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

func GetRoutesWithConditions(section string, dateFrom string, dateTo string, instructorId string, cntGroups string) (*dto.RouteIdsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllRouteIds(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetRoutesBySection, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, instructorId, dbqueries.GetRoutesByInstructor, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, cntGroups, dbqueries.GetRoutesByCntGroups, result)
	if err != nil {
		return nil, err
	}

	if dateTo != "" && dateFrom != "" {
		resultPart, err := dbqueries.GetRoutesByTime(pg, context.Background(), dateFrom, dateTo)
		if err != nil {
			return nil, err
		}
		result = intersection(result, resultPart)
	}

	var response dto.RouteIdsListResponse

	for _, routeId := range result {
		response.RouteIds = append(response.RouteIds, routeId.Id)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}

func GetRoutesWithGeoCond(placeId string, length string, difficulty string) (*dto.RouteIdsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllRouteIds(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, placeId, dbqueries.GetRoutesByPlace, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, length, dbqueries.GetRoutesByLength, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, difficulty, dbqueries.GetRoutesByDifficulty, result)
	if err != nil {
		return nil, err
	}

	var response dto.RouteIdsListResponse

	for _, routeId := range result {
		response.RouteIds = append(response.RouteIds, routeId.Id)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}

func GetInstructorsWithCondition(role string, routeType string, routeDifficulty string, cntTours string, tourId string, placeId string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllInstructors(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, role, dbqueries.GetInstructorsByRole, result)
	if err != nil {
		return nil, err
	}

	if routeType != "" && routeDifficulty != "" {

		typeReady, err := strconv.Atoi(routeType)
		if err != nil {
			return nil, err
		}
		difficultyReady, err := strconv.Atoi(routeDifficulty)
		if err != nil {
			return nil, err
		}

		resultPart, err := dbqueries.GetInstructorsByCategory(pg, context.Background(), typeReady, difficultyReady)
		if err != nil {
			return nil, err
		}

		result = intersection(result, resultPart)
	}

	result, err = checkParameter(pg, cntTours, dbqueries.GetInstructorsByCntTours, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, tourId, dbqueries.GetInstructorsByTour, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, placeId, dbqueries.GetInstructorsByPlace, result)
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

func GetTouristsWithTrainerInstructor(section string, group string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetTouristsWithTrainerInstructor(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetTouristsBySection, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, group, dbqueries.GetTouristsByGroup, result)

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

func GetTouristsCompletedALl(section string, group string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetTouristsCompletedAll(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetTouristsBySection, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, group, dbqueries.GetTouristsByGroup, result)

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

func GetTouristsCompletedRoutes(section string, group string, request dto.CompletedRoutesRequest) (*dto.PersonsListResponse, error) {
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

	for _, route := range request.Routes {
		resultPart, err := dbqueries.GetTouristsCompletedRoute(pg, context.Background(), int(route))

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

func GetAllRouteTypes() ([]dto.RouteType, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	typesM, err := dbqueries.GetAllRouteTypes(pg, context.Background())
	if err != nil {
		return nil, err
	}
	var response []dto.RouteType

	for _, types := range typesM {
		var jsonTypes dto.RouteType
		jsonTypes.Id = types.Id
		jsonTypes.Type = types.Type
		response = append(response, jsonTypes)
	}
	return response, nil
}

func GetSuitablePersonsByRoute(routeType string, difficulty string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	routeTypeInt, err := strconv.Atoi(routeType)
	if err != nil {
		return nil, err
	}

	var result []model.Person

	if routeTypeInt == 1 {
		result, err = dbqueries.GetTouristsBySection(pg, context.Background(), 2)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = dbqueries.GetAllTourists(pg, context.Background())
		if err != nil {
			return nil, err
		}
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
