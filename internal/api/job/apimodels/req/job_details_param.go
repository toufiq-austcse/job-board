package req

import "github.com/gin-gonic/gin"

type JobDetailsReqParam struct {
	Slug string `uri:"slug" binding:"required,min=1"`
}

func (model *JobDetailsReqParam) Validate(c *gin.Context) error {
	err := c.BindUri(&model)
	if err != nil {
		return err
	}
	return nil
}
