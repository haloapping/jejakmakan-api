package food

import (
	"github.com/jackc/pgx/v5"
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

func (r Repository) Add(c echo.Context, req AddReq) (Food, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	q := `
		INSERT INTO foods(id, user_id, owner_id, location_id images, name, description, price, review)
		VALUES(@id, @user_id, @owner_id, @location_id, @images, @name, @description, @price, @review)
		RETURNING *;
	`
	row := tx.QueryRow(
		ctx,
		q,
		pgx.NamedArgs{
			"id":          ulid.Make().String(),
			"user_id":     req.UserId,
			"owner_id":    req.OwnerId,
			"location_id": req.LocationId,
			"images":      req.Images,
			"name":        req.Name,
			"description": req.Description,
			"price":       req.Price,
			"review":      req.Review,
		},
	)
	var f Food
	err = row.Scan(
		&f.Id, &f.UserId, &f.OwnerId, &f.LocationId,
		&f.Images, &f.Name, &f.Description,
		&f.Price, &f.Review, &f.CreatedAt, &f.UpdatedAt,
	)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	q = `
		INSERT INTO food_stats(id, food_id)
		VALUES($1, $2);
	`
	_, err = tx.Exec(ctx, q, ulid.Make().String(), f.Id)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	tx.Commit(ctx)

	return f, nil
}
