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

		parameterReady, err := strconv.Atoi(parameter)

		if err != nil {
			return result, err
		}

		resultPart, err := searchFunc(pg, context.Background(), parameterReady)

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

func GetTrainersByWorkout(groupNum string, fromDate string, toDate string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	groupNumInt, err := strconv.Atoi(groupNum)
	if err != nil {
		return nil, err
	}

	if fromDate == "" {
		fromDate = "0001-01-01"
	}
	if toDate == "" {
		toDate = "2999-01-01"
	}

	result, err := dbqueries.GetTrainersByWorkout(pg, context.Background(), groupNumInt, fromDate, toDate)

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

func GetManagersWithCondition(salary string, birthYear string, age string, beginYear string, sex string) (*dto.PersonsListResponse, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := dbqueries.GetAllManagers(pg, context.Background())
	if err != nil {
		return nil, err
	}

	result, err = checkParameter(pg, salary, dbqueries.GetManagersBySalary, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, beginYear, dbqueries.GetManagersByBeginYear, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, birthYear, dbqueries.GetManagersByBirthYear, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, age, dbqueries.GetManagersByAge, result)
	if err != nil {
		return nil, err
	}
	result, err = checkParameter(pg, sex, dbqueries.GetManagersBySex, result)
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

func GetPersonRole(person string, section string) (*dto.PersonRole, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return nil, err
	}
	sectionInt, err := strconv.Atoi(section)
	if err != nil {
		return nil, err
	}

	role, err := dbqueries.GetPersonRole(pg, context.Background(), personInt, sectionInt)

	var jsonRole dto.PersonRole
	if role.Valid {
		jsonRole.Role = int(role.Int32)
		return &jsonRole, nil
	}
	return nil, nil
}

func SetPersonRole(person string, section string, role string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}
	sectionInt, err := strconv.Atoi(section)
	if err != nil {
		return err
	}
	roleInt, err := strconv.Atoi(role)
	if err != nil {
		return err
	}

	err = dbqueries.UpdatePersonRole(pg, context.Background(), personInt, sectionInt, roleInt)
	if err != nil {
		return err
	}
	return nil
}

func DeletePersonRole(person string, section string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}
	sectionInt, err := strconv.Atoi(section)
	if err != nil {
		return err
	}

	err = dbqueries.DeletePersonRole(pg, context.Background(), personInt, sectionInt)
	if err != nil {
		return err
	}
	return nil
}

func CreatePersonAttribute(attr dto.PersonAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.Attribute

	attrModel.Id = attr.Id
	attrModel.Name = attr.Name
	attrModel.Type = attr.Type
	if attr.Role >= 0 {
		attrModel.Role.Int32 = attr.Role
	} else {
		attrModel.Role.Valid = false
	}

	err = dbqueries.CreateAttribute(pg, context.Background(), attrModel)
	if err != nil {
		return err
	}
	return nil
}

func GetPersonAttribute(id string) (*dto.PersonAttribute, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	idInt, err := strconv.Atoi(id)

	var attr dto.PersonAttribute
	var attrModel *model.Attribute
	attrModel, err = dbqueries.GetAttribute(pg, context.Background(), idInt)
	if err != nil {
		return nil, err
	}
	if attrModel == nil {
		return nil, nil
	}
	attr.Name = attrModel.Name
	attr.Type = attrModel.Type
	attr.Id = attrModel.Id
	if attrModel.Role.Valid {
		attr.Role = int32(attrModel.Role.Int32)
	} else {
		attr.Role = -1
	}
	return &attr, nil
}

func SetPersonAttribute(attr dto.PersonAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.Attribute

	attrModel.Id = attr.Id
	attrModel.Name = attr.Name
	attrModel.Type = attr.Type
	if attr.Role >= 0 {
		attrModel.Role.Int32 = attr.Role
	} else {
		attrModel.Role.Valid = false
	}

	err = dbqueries.UpdateAttribute(pg, context.Background(), attrModel)
	if err != nil {
		return err
	}

	return nil
}

func DeletePersonAttribute(attr string) error {
	attrInt, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}
	err = dbqueries.DeleteAttribute(pg, context.Background(), attrInt)
	if err != nil {
		return err
	}

	return nil
}

func GetAllRoles() ([]dto.Role, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	rolesModel, err := dbqueries.GetAllRoles(pg, context.Background())
	if err != nil {
		return nil, err
	}

	var roles []dto.Role
	for _, role := range rolesModel {
		var jsonRole dto.Role
		jsonRole.Id = role.Id
		jsonRole.Role = role.Role
		roles = append(roles, jsonRole)
	}
	return roles, nil
}

func GetAllAttributes() ([]dto.PersonAttribute, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	attributesModel, err := dbqueries.GetAllAttributes(pg, context.Background())
	if err != nil {
		return nil, err
	}
	var attributes []dto.PersonAttribute
	for _, attr := range attributesModel {
		var jsonAttr dto.PersonAttribute
		jsonAttr.Id = attr.Id
		jsonAttr.Name = attr.Name
		jsonAttr.Type = attr.Type
		if attr.Role.Valid {
			jsonAttr.Role = int32(attr.Role.Int32)
		} else {
			jsonAttr.Role = -1
		}
		attributes = append(attributes, jsonAttr)
	}
	return attributes, nil
}

func GetPersonIntAttribute(person string, attribute string) (*int, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return nil, err
	}
	attributeInt, err := strconv.Atoi(attribute)
	if err != nil {
		return nil, err
	}

	val, err := dbqueries.GetPersonIntAttribute(pg, context.Background(), personInt, attributeInt)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func GetPersonFloatAttribute(person string, attribute string) (*float64, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return nil, err
	}
	attributeInt, err := strconv.Atoi(attribute)
	if err != nil {
		return nil, err
	}

	val, err := dbqueries.GetPersonFloatAttribute(pg, context.Background(), personInt, attributeInt)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func GetPersonStringAttribute(person string, attribute string) (*string, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return nil, err
	}
	attributeInt, err := strconv.Atoi(attribute)
	if err != nil {
		return nil, err
	}

	val, err := dbqueries.GetPersonStringAttribute(pg, context.Background(), personInt, attributeInt)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func GetPersonDateAttribute(person string, attribute string) (*string, error) {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return nil, err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return nil, err
	}
	attributeInt, err := strconv.Atoi(attribute)
	if err != nil {
		return nil, err
	}

	val, err := dbqueries.GetPersonDateAttribute(pg, context.Background(), personInt, attributeInt)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}

	t := val.Time
	stringVal := t.Format("2006-01-02")

	return &stringVal, nil
}

func SetPersonIntAttribute(attr dto.PersonIntAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.PersonIntAttribute
	attrModel.PersonId = attr.Person
	attrModel.AttributeId = attr.Attribute
	attrModel.Value = attr.Value

	val, err := dbqueries.GetPersonIntAttribute(pg, context.Background(), attr.Person, attr.Attribute)
	if err != nil {
		return err
	}

	if val == nil {
		err = dbqueries.SetPersonIntAttribute(pg, context.Background(), attrModel)

	} else {
		err = dbqueries.UpdatePersonIntAttribute(pg, context.Background(), attrModel)
	}

	if err != nil {
		return err
	}
	return nil
}

func SetPersonFloatAttribute(attr dto.PersonFloatAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.PersonFloatAttribute
	attrModel.PersonId = attr.Person
	attrModel.AttributeId = attr.Attribute
	attrModel.Value = attr.Value

	val, err := dbqueries.GetPersonFloatAttribute(pg, context.Background(), attr.Person, attr.Attribute)
	if err != nil {
		return err
	}

	if val == nil {
		err = dbqueries.SetPersonFloatAttribute(pg, context.Background(), attrModel)

	} else {
		err = dbqueries.UpdatePersonFloatAttribute(pg, context.Background(), attrModel)
	}

	if err != nil {
		return err
	}
	return nil
}

func SetPersonStringAttribute(attr dto.PersonStringAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.PersonStringAttribute
	attrModel.PersonId = attr.Person
	attrModel.AttributeId = attr.Attribute
	attrModel.Value = attr.Value

	val, err := dbqueries.GetPersonStringAttribute(pg, context.Background(), attr.Person, attr.Attribute)
	if err != nil {
		return err
	}

	if val == nil {
		err = dbqueries.SetPersonStringAttribute(pg, context.Background(), attrModel)

	} else {
		err = dbqueries.UpdatePersonStringAttribute(pg, context.Background(), attrModel)
	}

	if err != nil {
		return err
	}
	return nil
}

func SetPersonDateAttribute(attr dto.PersonDateAttribute) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	var attrModel model.PersonDateAttribute
	attrModel.PersonId = attr.Person
	attrModel.AttributeId = attr.Attribute
	err = attrModel.Value.Scan(attr.Value)
	if err != nil {
		return err
	}

	val, err := dbqueries.GetPersonDateAttribute(pg, context.Background(), attr.Person, attr.Attribute)
	if err != nil {
		return err
	}

	if val == nil {
		err = dbqueries.SetPersonDateAttribute(pg, context.Background(), attrModel)

	} else {
		err = dbqueries.UpdatePersonDateAttribute(pg, context.Background(), attrModel)
	}

	if err != nil {
		return err
	}
	return nil
}

func DeletePersonIntAttribute(person string, attr string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}

	attrInt, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}

	err = dbqueries.DeletePersonIntAttribute(pg, context.Background(), personInt, attrInt)
	if err != nil {
		return err
	}
	return nil
}

func DeletePersonFloatAttribute(person string, attr string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}

	attrInt, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}

	err = dbqueries.DeletePersonFloatAttribute(pg, context.Background(), personInt, attrInt)
	if err != nil {
		return err
	}
	return nil
}

func DeletePersonStringAttribute(person string, attr string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}

	attrInt, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}

	err = dbqueries.DeletePersonStringAttribute(pg, context.Background(), personInt, attrInt)
	if err != nil {
		return err
	}
	return nil
}

func DeletePersonDateAttribute(person string, attr string) error {
	pg, err := db.NewPG(context.Background())
	if err != nil {
		return err
	}

	personInt, err := strconv.Atoi(person)
	if err != nil {
		return err
	}

	attrInt, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}

	err = dbqueries.DeletePersonDateAttribute(pg, context.Background(), personInt, attrInt)
	if err != nil {
		return err
	}
	return nil
}
