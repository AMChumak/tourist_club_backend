package dto

type ChampionshipResponse struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type ChampionshipsListResponse struct {
	Page          int32                  `json:"page"`
	Total         int32                  `json:"total"`
	Championships []ChampionshipResponse `json:"championships"`
}
