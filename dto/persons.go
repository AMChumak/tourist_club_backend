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
