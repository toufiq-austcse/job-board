package req

import "github.com/gin-gonic/gin"

type LoginReqModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gt=5"`
}

func (model *LoginReqModel) Validate(c *gin.Context) error {
	err := c.BindJSON(model)
	if err != nil {
		return err
	}
	return nil
}
