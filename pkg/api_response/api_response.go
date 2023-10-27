package api_response

import "strings"

// swagger:parameters PaginationResponse
type PaginationResponse struct {
	TotalPage    int `json:"total_page"`
	ItemsPerPage int `json:"items_per_page"`
	CurrentPage  int `json:"current_page"`
	TotalItems   int `json:"total_items"`
}

// swagger:parameters Response
type Response struct {
	Code    int         `json:"code"` // This is Name
	Message string      `json:"message"`
	Errors  []string    `json:"handlers"`
	Data    interface{} `json:"data"`
}

// swagger:parameters ResponseWithPagination
type ResponseWithPagination struct {
	Code       int                 `json:"code"` // This is Name
	Message    string              `json:"message"`
	Errors     []string            `json:"handlers"`
	Data       interface{}         `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}
type EmptyObj struct {
}

func BuildResponse(code int, message string, data interface{}) Response {
	var errors = make([]string, 0)
	return Response{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}

func BuildResponseWithPagination(code int, message string, data interface{}, pagination *PaginationResponse) ResponseWithPagination {
	var errors = make([]string, 0)
	return ResponseWithPagination{
		Code:       code,
		Message:    message,
		Errors:     errors,
		Data:       data,
		Pagination: pagination,
	}
}

func BuildErrorResponse(code int, message string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")
	return Response{
		Code:    code,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
}
