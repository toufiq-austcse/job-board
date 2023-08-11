package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/service"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

func AuthMiddleware(companyService *companyService.CompanyService, ignoreAuthToken bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("auth middleware function called")
		authToken := context.GetHeader("Authorization")
		if authToken == "" {
			if ignoreAuthToken {
				context.Next()
				return
			}
			errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", "Authorization required in header", nil)
			context.AbortWithStatusJSON(errRes.Code, errRes)
			return
		}
		token, err := service.ValidateToken(authToken)
		if err != nil {
			errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
			context.AbortWithStatusJSON(errRes.Code, errRes)
			return
		}
		id, err := service.GetCompanyIdFromToken(token)
		if err != nil {
			errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
			context.AbortWithStatusJSON(errRes.Code, errRes)
			return
		}

		company, err := companyService.FindCompanyById(id, context)
		if err != nil {
			errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
			context.AbortWithStatusJSON(errRes.Code, errRes)
			return
		}
		context.Set("company", company)
		context.Next()
	}
}
