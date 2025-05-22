package dto

type PersonCreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type PersonResponse struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type PersonsListResponse struct {
	Page    int32            `json:"page"`
	Total   int32            `json:"total"`
	Persons []PersonResponse `json:"persons"`
}

type PersonRole struct {
	Role int `json:"role"`
}

type PersonAttribute struct {
	Id   int32  `json:"id"`
	Name string `json:"attr"`
	Role int32  `json:"role"`
	Type int32  `json:"attr_type"`
}

type PersonIntAttribute struct {
	Person    int `json:"person"`
	Attribute int `json:"attr"`
	Value     int `json:"value"`
}

type PersonFloatAttribute struct {
	Person    int     `json:"person"`
	Attribute int     `json:"attr"`
	Value     float64 `json:"value"`
}

type PersonStringAttribute struct {
	Person    int    `json:"person"`
	Attribute int    `json:"attr"`
	Value     string `json:"value"`
}

type PersonDateAttribute struct {
	Person    int    `json:"person"`
	Attribute int    `json:"attr"`
	Value     string `json:"value"`
}

type Group struct {
	Id          int32 `json:"id"`
	GroupNumber int32 `json:"group_number"`
	Section     int32 `json:"section"`
}

type Section struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
}

type Role struct {
	Id   int32  `json:"id"`
	Role string `json:"role"`
}
