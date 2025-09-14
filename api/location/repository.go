package location

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(p *pgxpool.Pool) Repository {
	return Repository{
		Pool: p,
	}
}

func (r Repository) Add(c echo.Context, req AddReq) (Location, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	q := `
		INSERT INTO locations(id, district, city, province, postal_code, details)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`
	row := tx.QueryRow(
		ctx,
		q,
		ulid.Make().String(), req.District, req.City, req.Province, req.PostalCode, req.Details,
	)
	c.Set("query", q)
	c.Set("queryArgs", req)
	var l Location
	err = row.Scan(&l.Id, &l.District, &l.City, &l.Province, &l.PostalCode, &l.Details, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}
	tx.Commit(ctx)

	return l, nil
}

func (r Repository) GetById(c echo.Context, id string) (Location, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	q := `
		SELECT *
		FROM locations
		WHERE id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	c.Set("query", q)
	c.Set("queryArgs", id)
	var l Location
	err = row.Scan(&l.Id, &l.District, &l.City, &l.Province, &l.PostalCode, &l.Details, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}
	tx.Commit(ctx)

	return l, nil
}

func (r Repository) GetAll(c echo.Context, limit int, offset int) ([]Location, int, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return []Location{}, 0, err
	}

	q := `
		SELECT COUNT(id)
		FROM (
			SELECT id
			FROM locations
		);
	`
	row := tx.QueryRow(ctx, q)
	c.Set("query", q)
	var total int
	err = row.Scan(&total)
	if err != nil {
		tx.Rollback(ctx)

		return []Location{}, 0, err
	}

	q = `
		SELECT COUNT(id)
		FROM (
			SELECT id
			FROM locations
			LIMIT $1
			OFFSET $2
		);
	`
	row = tx.QueryRow(ctx, q, limit, offset)
	c.Set("query", q)
	c.Set("queryArgs", []int{limit, offset})
	var count int
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback(ctx)

		return []Location{}, 0, err
	}

	q = `
		SELECT *
		FROM locations
		LIMIT $1
		OFFSET $2;
	`
	c.Set("query", q)
	c.Set("queryArgs", []int{limit, offset})
	rows, err := tx.Query(ctx, q, limit, offset)
	if err != nil {
		tx.Rollback(ctx)

		return []Location{}, 0, err
	}

	if limit > count {
		limit = count
	}
	idx := 0
	locations := make([]Location, count)
	for rows.Next() {
		var l Location
		err = rows.Scan(&l.Id, &l.District, &l.City, &l.Province, &l.PostalCode, &l.Details, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			tx.Rollback(ctx)

			return []Location{}, 0, err
		}
		locations[idx] = l
		idx++
	}

	fmt.Println(locations)

	return locations, total, nil
}

func (r Repository) UpdateById(c echo.Context, id string, req UpdateReq) (Location, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	argPos := 1
	columnArgs := make([]string, 0)
	columnValues := make([]interface{}, 0)
	if req.District == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("district = $%d", argPos))
		columnValues = append(columnValues, req.District)
		argPos++
	}
	if req.City == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("city = $%d", argPos))
		columnValues = append(columnValues, req.City)
		argPos++
	}
	if req.Province == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("province = $%d", argPos))
		columnValues = append(columnValues, req.Province)
		argPos++
	}
	if req.City == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("city = $%d", argPos))
		columnValues = append(columnValues, req.City)
		argPos++
	}
	if req.PostalCode == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("postalCode = $%d", argPos))
		columnValues = append(columnValues, req.PostalCode)
		argPos++
	}
	if req.Details == "" {
		columnArgs = append(columnArgs, fmt.Sprintf("details = $%d", argPos))
		columnValues = append(columnValues, req.Details)
		argPos++
	}

	columnValues = append(columnValues, id)

	q := fmt.Sprintf(`
			UPDATE locations
			SET %s
			WHERE id = $%d
			RETURNING *;
		`,
		strings.Join(columnArgs, ", "),
		argPos,
	)
	row := tx.QueryRow(ctx, q, columnValues...)
	c.Set("query", q)
	c.Set("queryArgs", []interface{}{id, req})
	var l Location
	err = row.Scan(&l.Id, &l.District, &l.City, &l.Province, &l.PostalCode, &l.Details, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	return l, nil
}

func (r Repository) DeleteById(c echo.Context, id string) (Location, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	q := `
		DELETE FROM locations
		WHERE id = $1
		RETURNING *;
	`
	c.Set("query", q)
	c.Set("queryArgs", id)
	row := tx.QueryRow(ctx, q, id)
	var l Location
	err = row.Scan(&l.Id, &l.District, &l.City, &l.Province, &l.PostalCode, &l.Details, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Location{}, err
	}

	return l, nil
}
