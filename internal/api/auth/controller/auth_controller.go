package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/apimodels/res"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// SignUp hosts godoc
// @Summary  Company SignUp
// @Param    request  body      req.SignUpReqModel  true  "Signup Req Body"
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/auth/signup [post]
// @Success  201      {object}  api_response.Response{data=res.SignUpResModel}
func (controller *AuthController) SignUp(context *gin.Context) {
	body := &req.SignUpReqModel{}
	if err := body.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, errRes)
		return
	}
	newCompany, token, expireAt, err := controller.authService.Signup(body, context)
	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	signUpRes := api_response.BuildResponse(http.StatusCreated, "Signup Successful", res.SignUpResModel{
		CompanyInfo: res.CompanyInfo{
			Name:  newCompany.Name,
			Email: newCompany.Email,
		},
		Token: res.Token{
			AccessToken: token,
			ExpireAt:    expireAt,
		},
	})
	context.JSON(signUpRes.Code, signUpRes)
}

// Login hosts godoc
// @Summary  Company Login
// @Param    request  body      req.LoginReqModel  true  "Login Req Body"
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/auth/login [post]
// @Success  201      {object}  api_response.Response{data=res.LoginResModel}
func (controller *AuthController) Login(context *gin.Context) {
	body := &req.LoginReqModel{}
	if err := body.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	company, token, expireAt, err := controller.authService.Login(body, context)
	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	loginRes := api_response.BuildResponse(http.StatusCreated, "Login Successful", res.LoginResModel{
		CompanyInfo: res.CompanyInfo{
			Name:  company.Name,
			Email: company.Email,
		},
		Token: res.Token{
			AccessToken: token,
			ExpireAt:    expireAt,
		},
	})
	context.JSON(loginRes.Code, loginRes)
}

// Me hosts godoc
// @Summary  Token Verification
// @Security Authorization
// @name Authorization
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/auth/me [get]
// @Success  200      {object}  api_response.Response{data=res.TokenVerificationRes}
func (controller *AuthController) Me(context *gin.Context) {
	company, _ := context.Get("company")
	entCompany := company.(*ent.Company)

	apiRes := api_response.BuildResponse(http.StatusOK, "", res.TokenVerificationRes{
		Name:        entCompany.Name,
		Location:    entCompany.Location,
		LogoUrl:     entCompany.LogoURL,
		WebsiteUrl:  entCompany.WebsiteURL,
		Email:       entCompany.Email,
		Slug:        entCompany.Slug,
		Size:        entCompany.Size,
		Industry:    entCompany.Industry,
		Established: entCompany.Established,
	})
	context.JSON(apiRes.Code, apiRes)
}
