package req

import "github.com/gin-gonic/gin"

type UpdateJobTaxonomyModel struct {
	Type string `json:"type"`
	Id   int    `json:"id"`
}
type UpdateJobReqModel struct {
	Title       string                   `json:"title,omitempty"`
	Taxonomies  []UpdateJobTaxonomyModel `json:"taxonomies,omitempty"`
	ApplyTo     string                   `json:"apply_to,omitempty"`
	Status      string                   `json:"status,omitempty"`
	Description string                   `json:"description,omitempty"`
}

func (model *UpdateJobReqModel) Validate(c *gin.Context) error {
	err := c.BindJSON(model)
	if err != nil {
		return err
	}
	return nil
}

type UpdateJobReqParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}

func (model *UpdateJobReqParam) Validate(c *gin.Context) error {
	err := c.BindUri(model)
	if err != nil {
		return err
	}
	return nil
}
