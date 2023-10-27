package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

func GlobalErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Internal Server Error", goErr.Error(), nil)
	c.AbortWithStatusJSON(500, errRes)
}
