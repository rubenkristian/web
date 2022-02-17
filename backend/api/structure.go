package api

type PaginationQuery struct {
	page int `form:"page"`
	size int `form:"size"`
}