package dto

type StrainResponse struct {
	Strain   string `json:"type"`
	Duration string `json:"duration"`
}

type StrainListResponse struct {
	Page       int32            `json:"page"`
	Total      int32            `json:"total"`
	StrainList []StrainResponse `json:"strainList"`
}
