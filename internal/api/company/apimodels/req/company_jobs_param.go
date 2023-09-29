package req

import "github.com/gin-gonic/gin"

type CompanyJobsReqParam struct {
	Slug string `uri:"slug" binding:"required,min=1"`
}

func (model *CompanyJobsReqParam) Validate(c *gin.Context) error {
	err := c.BindUri(&model)
	if err != nil {
		return err
	}
	return nil
}
