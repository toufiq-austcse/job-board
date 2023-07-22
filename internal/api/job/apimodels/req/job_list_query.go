package req

import "github.com/gin-gonic/gin"

type JobListQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (model *JobListQuery) Validate(c *gin.Context) error {
	err := c.Bind(&model)
	if err != nil {
		return err
	}
	return nil
}
