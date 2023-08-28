package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

type TaxonomyController struct {
	service *service.TaxonomyService
}

func NewTaxonomyController(service *service.TaxonomyService) *TaxonomyController {
	return &TaxonomyController{
		service: service,
	}
}

// ListTaxonomies hosts godoc
// @Summary  Get taxonomies
// @Param        type    query     string  false  "list taxonomies by category"
// @Tags     Taxonomy
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/taxonomies [get]
// @Success  201      {object}  api_response.Response{data=[]res.TaxonomyInListRes}
func (controller *TaxonomyController) ListTaxonomies(ctx *gin.Context) {
	taxonomyType, _ := ctx.GetQuery("type")

	taxonomies := controller.service.ListTaxonomy(taxonomyType, ctx)

	res := api_response.BuildResponse(http.StatusOK, "list taxonomies", taxonomies)
	ctx.JSON(res.Code, res)

}

// ListTaxonomyJobs hosts godoc
// @Summary  Get taxonomies
// @Param        type    query     string  false  "list taxonomies by category"
// @Tags     Taxonomy
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/taxonomies [get]
// @Success  201      {object}  api_response.Response{data=[]res.TaxonomyInListRes}
func (controller *TaxonomyController) ListTaxonomyJobs(ctx *gin.Context) {
	taxonomyType, _ := ctx.GetQuery("type")

	taxonomies := controller.service.ListTaxonomy(taxonomyType, ctx)

	res := api_response.BuildResponse(http.StatusOK, "list taxonomies", taxonomies)
	ctx.JSON(res.Code, res)

}
