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
	AttributeId int
	PersonId    int
	Value       string
}

type PersonIntAttribute struct {
	AttributeId int
	PersonId    int
	Value       int
}

type PersonFloatAttribute struct {
	AttributeId int
	PersonId    int
	Value       float64
}

type PersonDateAttribute struct {
	AttributeId int
	PersonId    int
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
