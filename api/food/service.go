package food

import "github.com/labstack/echo/v4"

type Service struct {
	Database
}

func NewService(db Database) Service {
	return Service{
		Database: db,
	}
}

func (s Service) Add(c echo.Context, req AddReq) (AddFood, error) {
	f, err := s.Database.Add(c, req)
	if err != nil {
		return AddFood{}, err
	}

	return f, nil
}

func (s Service) GetAll(c echo.Context, limit int, offset int) ([]Food, int, error) {
	foods, total, err := s.Database.GetAll(c, limit, offset)
	if err != nil {
		return []Food{}, 0, err
	}

	return foods, total, nil
}

func (s Service) GetById(c echo.Context, id string) (Food, error) {
	f, err := s.Database.GetById(c, id)
	if err != nil {
		return Food{}, err
	}

	return f, nil
}

func (s Service) UpdateById(c echo.Context, id string, req UpdateReq) (Food, error) {
	f, err := s.Database.UpdateById(c, id, req)
	if err != nil {
		return Food{}, err
	}

	return f, nil
}

func (s Service) DeleteById(c echo.Context, id string) (Food, error) {
	f, err := s.Database.DeleteById(c, id)
	if err != nil {
		return Food{}, err
	}

	return f, nil
}
