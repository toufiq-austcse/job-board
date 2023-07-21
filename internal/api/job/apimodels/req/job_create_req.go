package req

import "github.com/gin-gonic/gin"

type CreateJobReqModel struct {
	Title       string `json:"title" binding:"required"`
	Taxonomies  []int  `json:"taxonomies" binding:"required"`
	ApplyTo     string `json:"apply_to" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (model *CreateJobReqModel) Validate(c *gin.Context) error {
	err := c.BindJSON(model)
	if err != nil {
		return err
	}
	return nil
}
