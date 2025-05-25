package dto

type RouteIdsListResponse struct {
	Page     int32   `json:"page"`
	Total    int32   `json:"total"`
	RouteIds []int32 `json:"routeIds"`
}

type CompletedRoutesRequest struct {
	Routes []int32 `json:"routes"`
}

type RouteType struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}
