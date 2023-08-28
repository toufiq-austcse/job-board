package req

import "github.com/gin-gonic/gin"

type JobListQuery struct {
	Page         int    `form:"page,default=1" json:"page"`
	Limit        int    `form:"limit,default=50" json:"limit"`
	Status       string `form:"status" json:"status"`
	TaxonomySlug string `form:"taxonomy_slug" json:"taxonomy_slug"`
}

func (model *JobListQuery) Validate(c *gin.Context) error {
	err := c.Bind(&model)
	if err != nil {
		return err
	}
	return nil
}
