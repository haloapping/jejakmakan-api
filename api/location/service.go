package location

import "github.com/labstack/echo/v4"

type Service struct {
	Database
}

func NewService(db Database) Service {
	return Service{
		Database: db,
	}
}

func (s Service) Add(c echo.Context, req AddReq) (Location, error) {
	l, err := s.Database.Add(c, req)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}

func (s Service) GetAll(c echo.Context, limit int, offset int) ([]Location, int, error) {
	locations, total, err := s.Database.GetAll(c, limit, offset)
	if err != nil {
		return []Location{}, 0, err
	}

	return locations, total, nil
}

func (s Service) GetById(c echo.Context, id string) (Location, error) {
	l, err := s.Database.GetById(c, id)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}

func (s Service) UpdateById(c echo.Context, id string, req UpdateReq) (Location, error) {
	l, err := s.Database.UpdateById(c, id, req)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}

func (s Service) DeleteById(c echo.Context, id string) (Location, error) {
	l, err := s.Database.DeleteById(c, id)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}
