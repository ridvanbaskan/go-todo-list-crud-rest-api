package pagination

import (
	"net/http"
	"strconv"
)

type Pagination[T any] struct {
	TotalRecords int64 `json:"totalRecords"`
	PageSize     int   `json:"pageSize"`
	CurrentPage  int   `json:"currentPage"`
	TotalPages   int   `json:"totalPages"`
	Data         []T   `json:"data"`
}

func GetPaginationOffsetAndLimit(r *http.Request) (int, int, int) {
	pageQuery := r.URL.Query().Get("page")
	limitQuery := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageQuery != "" {
		pageInt, err := strconv.Atoi(pageQuery)
		if err == nil && pageInt > 0 {
			page = pageInt
		}
	}

	if limitQuery != "" {
		limitInt, err := strconv.Atoi(limitQuery)
		if err == nil && limitInt > 0 && limitInt <= 100 {
			limit = limitInt
		}
	}

	offset := (page - 1) * limit

	return page, offset, limit
}
