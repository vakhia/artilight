package dtos

type PaginationParams struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

type SortingParams struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}
