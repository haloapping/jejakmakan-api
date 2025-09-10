package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/haloapping/jejakmakan-api/config"
	"github.com/haloapping/jejakmakan-api/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	connStr := config.DBConnStr(".env")
	pool := db.NewConnection(connStr)
	defer pool.Close()

	nUser := flag.Int("nuser", 10, "number of user")
	nOwner := flag.Int("nowner", 10, "number of owner")
	nLocation := flag.Int("nlocation", 10, "number of location")
	nFood := flag.Int("nfood", 10, "number of food")
	flag.Parse()

	fmt.Println("Seed on progress... 🌱")
	err := GenerateFakeData(context.Background(), pool, *nUser, *nOwner, *nLocation, *nFood)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seed is done! 🌳")
}

func GenerateFakeData(ctx context.Context, p *pgxpool.Pool, nUser int, nOwner int, nLocation int, nFood int) error {
	userTx, err := p.Begin(ctx)
	if err != nil {
		userTx.Rollback(ctx)

		return err
	}

	userIdx := 0
	userIdxs := make([]string, nUser)
	for range nUser {
		q := `
		INSERT INTO users(id, profile_picture, username, email, fullname, password)
		VALUES(@id, @profile_picture, @username, @email, @fullname, @password)
		RETURNING id;
	`
		hashPassword, err := bcrypt.GenerateFromPassword([]byte("Secretp4ssword!"), bcrypt.DefaultCost)
		if err != nil {
			userTx.Rollback(ctx)

			return err
		}
		row := userTx.QueryRow(
			ctx,
			q,
			pgx.NamedArgs{
				"id":              ulid.Make().String(),
				"profile_picture": gofakeit.Word() + ".jpg",
				"username":        gofakeit.Username(),
				"email":           gofakeit.Email(),
				"fullname":        gofakeit.Name(),
				"password":        string(hashPassword),
			},
		)
		var id string
		err = row.Scan(&id)
		if err != nil {
			userTx.Rollback(ctx)

			return err
		}
		userIdxs[userIdx] = id
		userIdx++
	}
	userTx.Commit(ctx)
	fmt.Println("User seed is done!")

	ownerTx, err := p.Begin(ctx)
	if err != nil {
		ownerTx.Rollback(ctx)

		return err
	}
	ownerIdx := 0
	ownerIdxs := make([]string, nUser)
	for range nOwner {
		q := `
			INSERT INTO owners(id, images, name)
			VALUES(@id, @images, @name)
			RETURNING id;
		`
		row := ownerTx.QueryRow(
			ctx,
			q,
			pgx.NamedArgs{
				"id":     ulid.Make().String(),
				"images": gofakeit.Word() + ".jpg",
				"name":   gofakeit.Username(),
			},
		)
		var id string
		err = row.Scan(&id)
		if err != nil {
			ownerTx.Rollback(ctx)

			return err
		}
		ownerIdxs[ownerIdx] = id
		ownerIdx++
	}
	ownerTx.Commit(ctx)
	fmt.Println("Owner seed is done!")

	locationTx, err := p.Begin(ctx)
	if err != nil {
		locationTx.Rollback(ctx)

		return err
	}
	locationIdx := 0
	locationIdxs := make([]string, nUser)
	for range nOwner {
		q := `
			INSERT INTO locations(id, district, city, province, postal_code, details)
			VALUES(@id, @district, @city, @province, @postal_code, @details)
			RETURNING id;
		`
		row := locationTx.QueryRow(
			ctx,
			q,
			pgx.NamedArgs{
				"id":          ulid.Make().String(),
				"district":    gofakeit.Address().State,
				"city":        gofakeit.Address().City,
				"province":    gofakeit.Address().Unit,
				"postal_code": gofakeit.Address().Zip,
				"details":     gofakeit.Address().Street,
			},
		)
		var id string
		err = row.Scan(&id)
		if err != nil {
			locationTx.Rollback(ctx)

			return err
		}
		locationIdxs[locationIdx] = id
		locationIdx++
	}
	locationTx.Commit(ctx)
	fmt.Println("Location seed is done!")

	foodTx, err := p.Begin(ctx)
	if err != nil {
		foodTx.Rollback(ctx)

		return err
	}
	foodIdx := 0
	foodIdxs := make([]string, nFood)
	for range nFood {
		q := `
			INSERT INTO foods(id, user_id, owner_id, location_id, images, name, description, price, review)
			VALUES(@id, @user_id, @owner_id, @location_id, @images, @name, @description, @price, @review)
			RETURNING id;
		`
		randomUser := userIdxs[rand.IntN(nUser)]
		randomOwner := ownerIdxs[rand.IntN(nOwner)]
		randomLocation := locationIdxs[rand.IntN(nLocation)]
		row := foodTx.QueryRow(
			ctx,
			q,
			pgx.NamedArgs{
				"id":          ulid.Make().String(),
				"user_id":     randomUser,
				"owner_id":    randomOwner,
				"location_id": randomLocation,
				"images":      gofakeit.Word() + ".jpg",
				"name":        gofakeit.Name(),
				"description": gofakeit.Sentence(10),
				"price":       rand.IntN(150_000),
				"review":      gofakeit.Sentence(20),
			},
		)
		var id string
		err = row.Scan(&id)
		if err != nil {
			foodTx.Rollback(ctx)

			return err
		}
		foodIdxs[foodIdx] = id
		foodIdx++
	}
	foodTx.Commit(ctx)
	fmt.Println("Food seed is done!")

	foodStatTx, err := p.Begin(ctx)
	if err != nil {
		foodStatTx.Rollback(ctx)

		return err
	}
	for i := range nFood {
		q := `
			INSERT INTO food_stats(id, food_id)
			VALUES($1, $2)
		`
		_, err = foodStatTx.Exec(ctx, q, ulid.Make().String(), foodIdxs[i])
		if err != nil {
			foodTx.Rollback(ctx)

			return err
		}
	}
	foodStatTx.Commit(ctx)
	fmt.Println("Food Stat seed is done!")

	return nil
}
