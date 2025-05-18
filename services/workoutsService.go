package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
	"strconv"
)

func GetStrainForTrainer(trainer string, fromDate string, toDate string) (*dto.StrainListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	if fromDate == "" {
		fromDate = "0001-01-01"
	}
	if toDate == "" {
		toDate = "2999-01-01"
	}

	idReady, err := strconv.Atoi(trainer)

	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetStrainForTrainer(pg, context.Background(), idReady, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	var response dto.StrainListResponse

	for _, strain := range result {
		var jsonStrain dto.StrainResponse
		jsonStrain.Strain = strain.Type
		jsonStrain.Duration = strain.GetTimeAsString()

		response.StrainList = append(response.StrainList, jsonStrain)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}
