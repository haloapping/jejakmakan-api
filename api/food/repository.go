package food

import (
	"fmt"
	"strings"

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

func (r Repository) Add(c echo.Context, req AddReq) (AddFood, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return AddFood{}, err
	}

	q := `
		INSERT INTO foods(id, user_id, owner_id, location_id,images, name, description, price, review)
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
	c.Set("query", q)
	c.Set("queryArgs", req)
	var f AddFood
	err = row.Scan(
		&f.Id, &f.UserId, &f.OwnerId, &f.LocationId,
		&f.Images, &f.Name, &f.Description,
		&f.Price, &f.Review, &f.CreatedAt, &f.UpdatedAt,
	)
	if err != nil {
		tx.Rollback(ctx)

		return AddFood{}, err
	}

	q = `
		INSERT INTO food_stats(id, food_id)
		VALUES($1, $2);
	`
	_, err = tx.Exec(ctx, q, ulid.Make().String(), f.Id)
	if err != nil {
		tx.Rollback(ctx)

		return AddFood{}, err
	}

	tx.Commit(ctx)

	return f, nil
}

func (r Repository) GetAll(c echo.Context, limit int, offset int) ([]Food, int, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return []Food{}, 0, err
	}

	q := `
		SELECT COUNT(id)
		FROM (
			SELECT (id)
			FROM foods
		);
	`
	row := tx.QueryRow(ctx, q)
	var total int
	err = row.Scan(&total)
	if err != nil {
		tx.Rollback(ctx)

		return []Food{}, 0, err
	}

	q = `
		SELECT COUNT(id)
		FROM (
			SELECT id
			FROM foods
			LIMIT $1
			OFFSET $2
		);
	`
	row = tx.QueryRow(ctx, q, limit, offset)
	var count int
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback(ctx)

		return []Food{}, 0, err
	}

	q = `
		SELECT f.id, u.username, o.name, l.district, f.images, f.name, f.description, f.price, f.review, f.created_at, f.updated_at
		FROM foods f
		JOIN users u ON f.user_id = u.id
		JOIN owners o ON f.owner_id = o.id
		JOIN locations l ON f.location_id = l.id
		LIMIT $1
		OFFSET $2;
	`
	rows, err := tx.Query(ctx, q, limit, offset)
	if err != nil {
		tx.Rollback(ctx)

		return []Food{}, 0, err
	}

	idx := 0
	foods := make([]Food, count)
	for rows.Next() {
		var f Food
		err := rows.Scan(
			&f.Id,
			&f.Username,
			&f.Ownername,
			&f.Location,
			&f.Images,
			&f.Name,
			&f.Description,
			&f.Price,
			&f.Review,
			&f.CreatedAt,
			&f.UpdatedAt,
		)
		if err != nil {
			tx.Rollback(ctx)

			return []Food{}, 0, err
		}
		foods[idx] = f
		idx++
	}
	tx.Commit(ctx)

	return foods, total, nil
}

func (r Repository) GetById(c echo.Context, id string) (Food, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	q := `
		SELECT f.id, u.username, o.name, l.district, f.images, f.name, f.description, f.price, f.review, f.created_at, f.updated_at
		FROM foods f
		JOIN users u ON f.user_id = u.id
		JOIN owners o ON f.owner_id = o.id
		JOIN locations l ON f.location_id = l.id
		WHERE f.id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	var f Food
	err = row.Scan(
		&f.Id,
		&f.Username,
		&f.Ownername,
		&f.Location,
		&f.Images,
		&f.Name,
		&f.Description,
		&f.Price,
		&f.Review,
		&f.CreatedAt,
		&f.UpdatedAt,
	)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}
	tx.Commit(ctx)

	return f, nil
}

func (r Repository) UpdateById(c echo.Context, id string, req UpdateReq) (Food, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	columnNames := make([]string, 0)
	columnArgs := pgx.NamedArgs{}
	if req.UserId != "" {
		columnNames = append(columnNames, "user_id = @user_id")
		columnArgs["user_id"] = req.UserId
	}
	if req.OwnerId != "" {
		columnNames = append(columnNames, "owner_id = @owner_id")
		columnArgs["owner_id"] = req.OwnerId
	}
	if req.LocationId != "" {
		columnNames = append(columnNames, "location_id = @location_id")
		columnArgs["location_id"] = req.LocationId
	}
	if req.Images != "" {
		columnNames = append(columnNames, "images = @images")
		columnArgs["images"] = req.Images
	}
	if req.Name != "" {
		columnNames = append(columnNames, "name = @name")
		columnArgs["name"] = req.Name
	}
	if req.Description != "" {
		columnNames = append(columnNames, "description = @description")
		columnArgs["description"] = req.Description
	}
	if req.Price != 0 {
		columnNames = append(columnNames, "price = @price")
		columnArgs["price"] = req.Price
	}
	if req.Review != "" {
		columnNames = append(columnNames, "review = @review")
		columnArgs["review"] = req.Review
	}
	columnArgs["id"] = id
	q := fmt.Sprintf(`
		UPDATE foods f
		SET %s
		FROM users u, owners o, locations l
		WHERE
			f.id = @id AND
			f.user_id = u.id AND
			f.owner_id = o.id AND
			f.location_id = l.id
		RETURNING f.id, u.username, o.name, l.district, f.images, f.name, f.description, f.price, f.review, f.created_at, f.updated_at;
	`, strings.Join(columnNames, ", "))
	fmt.Println("Query:", q)
	fmt.Println("Args:", columnArgs)
	c.Set("query", q)
	c.Set("queryArgs", columnArgs)
	row := tx.QueryRow(ctx, q, columnArgs)
	var f Food
	err = row.Scan(
		&f.Id,
		&f.Username,
		&f.Ownername,
		&f.Location,
		&f.Images,
		&f.Name,
		&f.Description,
		&f.Price,
		&f.Review,
		&f.CreatedAt,
		&f.UpdatedAt,
	)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}
	tx.Commit(ctx)

	return f, nil
}

func (r Repository) DeleteById(c echo.Context, id string) (Food, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}

	q := `
		DELETE FROM foods f
		USING users u, owners o, locations l
		WHERE
			f.id = $1 AND
			f.user_id = u.id AND
			f.owner_id = o.id AND
			f.location_id = l.id
		RETURNING f.id, u.username, o.name, l.district, f.images, f.name, f.description, f.price, f.review, f.created_at, f.updated_at;
	`
	row := tx.QueryRow(ctx, q, id)
	var f Food
	err = row.Scan(
		&f.Id,
		&f.Username,
		&f.Ownername,
		&f.Location,
		&f.Images,
		&f.Name,
		&f.Description,
		&f.Price,
		&f.Review,
		&f.CreatedAt,
		&f.UpdatedAt,
	)
	if err != nil {
		tx.Rollback(ctx)

		return Food{}, err
	}
	tx.Commit(ctx)

	return f, nil
}
