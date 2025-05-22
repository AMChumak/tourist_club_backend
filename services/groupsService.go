package services

import (
	"context"
	"db_backend/db"
	"db_backend/dbqueries"
	"db_backend/dto"
	"db_backend/model"
	"strconv"
)

func CreateGroup(group dto.Group) (int, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return -1, err
	}

	var groupModel model.Group

	groupModel.Id = group.Id
	groupModel.GroupNumber = group.GroupNumber
	groupModel.Section = group.Section

	newId, err := dbqueries.CreateGroup(pg, context.Background(), groupModel)
	if err != nil {
		return -1, err
	}
	return newId, nil
}

func GetGroup(id string) (*model.Group, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	group, err := dbqueries.GetGroup(pg, context.Background(), idInt)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func UpdateGroup(group dto.Group) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var groupModel model.Group
	groupModel.Id = group.Id
	groupModel.GroupNumber = group.GroupNumber
	groupModel.Section = group.Section

	err = dbqueries.UpdateGroup(pg, context.Background(), groupModel)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGroup(id string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = dbqueries.DeleteGroup(pg, context.Background(), idInt)
	if err != nil {
		return err
	}
	return nil
}

func GetGroupMembers(group string) ([]dto.PersonResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}
	groupIdInt, err := strconv.Atoi(group)
	if err != nil {
		return nil, err
	}
	persons, err := dbqueries.GetGroupMembers(pg, context.Background(), groupIdInt)
	if err != nil {
		return nil, err
	}

	var members []model.Person

	for _, person := range persons {
		member, err := dbqueries.GetPerson(pg, context.Background(), person)
		if err != nil {
			return nil, err
		}
		members = append(members, *member)
	}

	var result []dto.PersonResponse
	for _, person := range members {
		var personResponse dto.PersonResponse
		personResponse.Id = person.Id
		personResponse.Name = person.Name
		personResponse.Surname = person.Surname
		personResponse.Patronymic = person.Patronymic
		result = append(result, personResponse)
	}
	return result, nil
}

func AddGroupMember(group string, person string) error {

	groupIdInt, err := strconv.Atoi(group)
	if err != nil {
		return err
	}
	personIdInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}

	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}
	err = dbqueries.AddGroupMember(pg, context.Background(), personIdInt, groupIdInt)
	if err != nil {
		return err
	}
	return nil
}

func RemoveGroupMember(group string, person string) error {
	groupIdInt, err := strconv.Atoi(group)
	if err != nil {
		return err
	}
	personIdInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}
	err = dbqueries.RemoveGroupMember(pg, context.Background(), personIdInt, groupIdInt)
	if err != nil {
		return err
	}
	return nil
}

func GetAllGroups() ([]dto.Group, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	groups, err := dbqueries.GetGroups(pg, context.Background())
	if err != nil {
		return nil, err
	}

	var result []dto.Group
	for _, group := range groups {
		var groupResponse dto.Group
		groupResponse.Id = group.Id
		groupResponse.GroupNumber = group.GroupNumber
		groupResponse.Section = group.Section
		result = append(result, groupResponse)
	}

	return result, nil
}

func GetAllSections() ([]dto.Section, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	sections, err := dbqueries.GetAllSections(pg, context.Background())
	if err != nil {
		return nil, err
	}

	var result []dto.Section
	for _, section := range sections {
		var jsonSection dto.Section
		jsonSection.Id = section.Id
		jsonSection.Title = section.Title
		result = append(result, jsonSection)
	}

	return result, nil
}
