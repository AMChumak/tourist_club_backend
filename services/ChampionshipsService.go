package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
)

func GetChampionshipsWithCondition(section string) (*dto.ChampionshipsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllChampionships(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, section, dbqueries.GetAllChampionshipsBySection, result)
	if err != nil {
		return nil, err
	}

	var response dto.ChampionshipsListResponse

	for _, championship := range result {
		var jsonChamp dto.ChampionshipResponse
		jsonChamp.Id = championship.Id
		jsonChamp.Title = championship.Title
		jsonChamp.Date = championship.GetDateAsString()

		response.Championships = append(response.Championships, jsonChamp)
	}

	response.Total = 1
	response.Page = 0

	return &response, nil
}
