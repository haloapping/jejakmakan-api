package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(p *pgxpool.Pool) Database {
	return Database{
		Pool: p,
	}
}

func (db Database) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	ctx := c.Request().Context()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserRegister{}, err
	}

	q := `
		INSERT INTO users(id, profile_picture, username, email, password, fullname)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, profile_picture, username, email, fullname, created_at, updated_at;
	`
	row := tx.QueryRow(
		ctx,
		q,
		ulid.Make().String(), req.ProfilePicture, req.Username, req.Email, req.Password, req.Fullname,
	)
	var ur UserRegister
	err = row.Scan(&ur.Id, &ur.ProfilePicture, &ur.Username, &ur.Email, &ur.Fullname, &ur.CreatedAt, &ur.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return UserRegister{}, err
	}
	tx.Commit(ctx)

	return ur, nil
}

func (db Database) Login(c echo.Context, req UserLoginReq) (UserLogin, error) {
	ctx := c.Request().Context()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserLogin{}, err
	}

	q := `
		SELECT id, username, password
		FROM users
		WHERE username = $1;
	`
	row := tx.QueryRow(ctx, q, req.Username)
	c.Set("query", q)
	c.Set("queryArgs", req.Username)
	var ul UserLogin
	err = row.Scan(&ul.Id, &ul.Username, &ul.Password)
	if err != nil {
		tx.Rollback(ctx)

		return UserLogin{}, err
	}
	tx.Commit(ctx)

	return ul, err
}

func (db Database) Biodata(c echo.Context, username string) (UserBiodata, error) {
	ctx := c.Request().Context()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}

	q := `
		SELECT id, profile_picture, username, email, fullname, password
		FROM users
		WHERE username = $1;
	`
	row := tx.QueryRow(ctx, q, username)
	c.Set("query", q)
	c.Set("queryArgs", username)
	var ub UserBiodata
	err = row.Scan(&ub.Id, &ub.ProfilePicture, &ub.Username, &ub.Password, &ub.Email, &ub.Fullname)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}
	tx.Commit(ctx)

	return UserBiodata{}, nil
}
