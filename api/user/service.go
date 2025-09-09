package user

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s Service) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserRegister{}, err
	}

	req.Password = string(hashPassword)
	items, err := s.Repository.Register(c, req)
	if err != nil {
		return UserRegister{}, err
	}

	return items, nil
}

func (s Service) Login(c echo.Context, req UserLoginReq) (UserLogin, error) {
	items, err := s.Repository.Login(c, req)
	if err != nil {
		return UserLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(items.Password), []byte(req.Password))
	if err != nil {
		return UserLogin{}, err
	}

	return items, nil
}

func (s Service) Biodata(c echo.Context, username string) (UserBiodata, error) {
	items, err := s.Repository.Biodata(c, username)
	if err != nil {
		return UserBiodata{}, err
	}

	return items, nil
}
