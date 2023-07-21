package controller

import "github.com/gin-gonic/gin"

type JobController struct {
}

func NewJobController() *JobController {
	return &JobController{}
}

// Create hosts godoc
// @Summary  Create New Job
// @Param    request  body      req.SignUpReqModel  true  "Signup Req Body"
// @Tags     Jobs
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/jobs [post]
// @Success  201      {object}  api_response.Response{data=res.SignUpResModel}
func (controller *JobController) Create(context *gin.Context) {

}
