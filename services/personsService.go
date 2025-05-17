package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
	"db_backend/model"
)

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
