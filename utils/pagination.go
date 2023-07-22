package utils

import (
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"math"
)

func GetPaginationData(total int, currentPage int, limit int) *api_response.PaginationResponse {
	totalPage := 0
	if limit > 0 {
		totalPage = int(math.Ceil(float64(total) / float64(limit)))
	}
	return &api_response.PaginationResponse{
		TotalPage:    totalPage,
		ItemsPerPage: limit,
		CurrentPage:  currentPage,
		TotalItems:   total,
	}

}
