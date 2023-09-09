package req

import "github.com/gin-gonic/gin"

type CompanyDetailsReqParam struct {
	Slug string `uri:"slug" binding:"required,min=1"`
}

func (model *CompanyDetailsReqParam) Validate(c *gin.Context) error {
	err := c.BindUri(&model)
	if err != nil {
		return err
	}
	return nil
}
