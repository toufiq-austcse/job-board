package service

import (
	"context"
	"errors"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/apimodels/req"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	companyService *companyService.CompanyService
}

func NewAuthService(companyService *companyService.CompanyService) *AuthService {
	return &AuthService{
		companyService: companyService,
	}
}

func (authService *AuthService) Signup(dto *req.SignUpReqModel, context context.Context) (company *ent.Company, token string, expireAt int64, signUpError error) {
	newCompany, err := authService.companyService.CreateCompany(dto.Name, dto.Email, dto.Password, context)
	if err != nil {
		return nil, "", 0, err
	}
	newToken, tokenExpireAt, err := GenerateToken(newCompany.ID)
	if err != nil {
		return nil, "", 0, err
	}
	return newCompany, newToken, tokenExpireAt, nil

}
func (authService *AuthService) Login(dto *req.LoginReqModel, context context.Context) (company *ent.Company, token string, expireAt int64, loginError error) {
	existingCompany, err := authService.companyService.FindCompanyByEmail(dto.Email, context)
	if err != nil {
		return nil, "", 0, err
	}
	isPasswordMatched, err := checkPassword(existingCompany.Password, dto.Password)
	if err != nil {
		return nil, "", 0, err
	}
	if !isPasswordMatched {
		return nil, "", 0, errors.New("Incorrect password")
	}
	newToken, tokenExpireAt, err := GenerateToken(existingCompany.ID)
	if err != nil {
		return nil, "", 0, err
	}
	return existingCompany, newToken, tokenExpireAt, nil

}

func checkPassword(hashedPassword string, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
