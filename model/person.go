package model

import "github.com/jackc/pgx/v5/pgtype"

type Person struct {
	Id         int32
	Name       string
	Surname    string
	Patronymic string
}

type Attribute struct {
	Id   int32
	Name string
	Type int32
	Role pgtype.Int4
}

type PersonStringAttribute struct {
	AttributeId int32
	PersonId    int32
	Value       string
}

type PersonIntAttribute struct {
	AttributeId int32
	PersonId    int32
	Value       int32
}

type PersonFloatAttribute struct {
	AttributeId int32
	PersonId    int32
	Value       float64
}

type PersonDateAttribute struct {
	AttributeId int32
	PersonId    int32
	Value       pgtype.Date
}

type Group struct {
	Id          int32
	GroupNumber int32
	Section     int32
}

type Section struct {
	Id    int32
	Title string
}

type Role struct {
	Id   int32
	Role string
}
