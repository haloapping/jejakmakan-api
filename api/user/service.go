package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Database
}

func NewService(db Database) Service {
	return Service{
		Database: db,
	}
}

func (s Service) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserRegister{}, err
	}

	req.Password = string(hashPassword)
	items, err := s.Database.Register(c, req)
	if err != nil {
		return UserRegister{}, err
	}

	return items, nil
}

func (s Service) Login(c echo.Context, req UserLoginReq) (UserLogin, error) {
	items, err := s.Database.Login(c, req)
	if err != nil {
		return UserLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(items.Password), []byte(req.Password))
	if err != nil {
		return UserLogin{}, fmt.Errorf("invalid username or password")
	}

	return items, nil
}

func (s Service) Biodata(c echo.Context, username string) (UserBiodata, error) {
	items, err := s.Database.Biodata(c, username)
	if err != nil {
		return UserBiodata{}, err
	}

	return items, nil
}
