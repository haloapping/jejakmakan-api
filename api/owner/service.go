package owner

import (
	"github.com/labstack/echo/v4"
)

type Service struct {
	Database
}

func NewService(db Database) Service {
	return Service{
		Database: db,
	}
}

func (s Service) Add(c echo.Context, req AddReq) (Owner, error) {
	o, err := s.Database.Add(c, req)
	if err != nil {
		return Owner{}, err
	}

	return o, nil
}

func (s Service) GetAll(c echo.Context, limit int, offset int) ([]Owner, int, error) {
	owners, total, err := s.Database.GetAll(c, limit, offset)
	if err != nil {
		return []Owner{}, 0, err
	}

	return owners, total, nil
}

func (s Service) GetById(c echo.Context, id string) (Owner, error) {
	o, err := s.Database.GetById(c, id)
	if err != nil {
		return Owner{}, err
	}

	return o, nil
}

func (s Service) UpdateById(c echo.Context, id string, req UpdateReq) (Owner, error) {
	o, err := s.Database.UpdateById(c, id, req)
	if err != nil {
		return Owner{}, err
	}

	return o, nil
}

func (s Service) DeleteById(c echo.Context, id string) (Owner, error) {
	o, err := s.Database.DeleteById(c, id)
	if err != nil {
		return Owner{}, err
	}

	return o, nil
}
