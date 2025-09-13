package owner

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

func (r Repository) Add(c echo.Context, req AddReq) (Owner, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}

	q := `
		INSERT INTO owners(id, images, name)
		VALUES($1, $2, $3)
		RETURNING *;
	`
	row := tx.QueryRow(
		ctx,
		q,
		ulid.Make().String(), req.Images, req.Name,
	)
	var o Owner
	err = row.Scan(&o.Id, &o.Images, &o.Name, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}
	tx.Commit(ctx)

	return o, nil
}

func (r Repository) GetAll(c echo.Context, limit int, offset int) ([]Owner, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return []Owner{}, err
	}

	q := `
		SELECT COUNT(id)
		FROM (
			SELECT id
			FROM owners
			LIMIT $1
			OFFSET $2
		);
	`
	row := tx.QueryRow(ctx, q, limit, offset)
	var count int
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback(ctx)

		return []Owner{}, err
	}

	q = `
		SELECT *
		FROM owners
		LIMIT $1
		OFFSET $2;
	`
	rows, err := tx.Query(ctx, q, limit, offset)
	if err != nil {
		tx.Rollback(ctx)

		return []Owner{}, err
	}
	if limit > count {
		limit = count
	}
	idx := 0
	owners := make([]Owner, count)
	for rows.Next() {
		var o Owner
		err = rows.Scan(&o.Id, &o.Images, &o.Name, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			tx.Rollback(ctx)

			return []Owner{}, err
		}
		owners[idx] = o
		idx++
	}

	return owners, nil
}

func (r Repository) GetById(c echo.Context, id string) (Owner, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}

	q := `
		SELECT *
		FROM owners
		WHERE id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	var o Owner
	err = row.Scan(&o.Id, &o.Images, &o.Name, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}

	return o, nil
}

func (r Repository) UpdateById(c echo.Context, id string, req UpdateReq) (Owner, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}

	argPos := 1
	columnArgs := make([]string, 0)
	columnValues := make([]interface{}, 0)
	if req.Images != "" {
		columnArgs = append(columnArgs, fmt.Sprintf("images = $%d", argPos))
		columnValues = append(columnValues, req.Images)
		argPos++
	}
	if req.Name != "" {
		columnArgs = append(columnArgs, fmt.Sprintf("name = $%d", argPos))
		columnValues = append(columnValues, req.Images)
		argPos++
	}

	columnValues = append(columnValues, id)

	q := fmt.Sprintf(`
		UPDATE owners
		SET $%s
		WHERE id $%d;
	`,
		strings.Join(columnArgs, ", "),
		argPos,
	)
	row := tx.QueryRow(ctx, q, columnValues...)
	var o Owner
	err = row.Scan(&o.Id, &o.Images, &o.Name, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}
	tx.Commit(ctx)

	return o, nil
}

func (r Repository) DeleteById(c echo.Context, id string) (Owner, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}

	q := `
		DELETE FROM owners
		WHERE id = $1
		RETURNING *;
	`
	row := tx.QueryRow(ctx, q, id)
	var o Owner
	err = row.Scan(&o.Id, &o.Images, &o.Name, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Owner{}, err
	}
	tx.Commit(ctx)

	return o, nil
}
