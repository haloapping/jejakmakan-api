package user

import (
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

func (r Repository) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserRegister{}, err
	}

	q := `
		INSERT INTO users(id, profile_picture, username, email, fullname)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, profile_picture, username, email, fullname, created_at, updated_at;
	`
	row := tx.QueryRow(
		ctx,
		q,
		ulid.Make().String(), req.ProfilePicture, req.Username, req.Email, req.Fullname,
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

func (r Repository) Login(c echo.Context, req UserLoginReq) (UserLogin, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
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
	var ul UserLogin
	err = row.Scan(&ul.Id, &ul.Username, &ul.Password)
	if err != nil {
		tx.Rollback(ctx)

		return UserLogin{}, err
	}
	tx.Commit(ctx)

	return ul, err
}

func (r Repository) Biodata(c echo.Context, username string) (UserBiodata, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}

	q := `
		SELECT id, profile_picture, username, password, email, fullname
		FROM users
		WHERE username = $1;
	`
	row := tx.QueryRow(ctx, q, username)
	var ub UserBiodata
	err = row.Scan(&ub.Id, &ub.ProfilePicture, &ub.Username, &ub.Password, &ub.Email, &ub.Fullname)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}
	tx.Commit(ctx)

	return UserBiodata{}, nil
}
