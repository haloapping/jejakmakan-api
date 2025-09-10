package food

import "github.com/labstack/echo/v4"

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s Service) Add(c echo.Context, req AddReq) (Food, error) {
	f, err := s.Repository.Add(c, req)
	if err != nil {
		return Food{}, err
	}

	return f, nil
}
