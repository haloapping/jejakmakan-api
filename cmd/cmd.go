//go:build cmd

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/haloapping/jejakmakan-api/db"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	migrationDir = "../migrations"
	dbDriver     = "postgres"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("prod.config")
	err := viper.ReadInConfig()
	if err != nil {
		zlog.Error().Msg(err.Error())
	}
	viper.AutomaticEnv()

	dbUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command>")
		return
	}
	switch os.Args[1] {
	case "dbup":
		cmdDBUp := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "up")
		fmt.Println("Running command:", cmdDBUp.String())
		if err := cmdDBUp.Run(); err != nil {
			fmt.Println("Error running goose up:", err)
			return
		}
	case "dbdown":
		cmdDBDown := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "down")
		fmt.Println("Running command:", cmdDBDown.String())
		if err := cmdDBDown.Run(); err != nil {
			fmt.Println("Error running goose down:", err)
			return
		}
	case "dbreset":
		cmdDBReset := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "down")
		fmt.Println("Running command:", cmdDBReset.String())
		if err := cmdDBReset.Run(); err != nil {
			fmt.Println("Error running goose down:", err)
			return
		}

		cmdDBUp := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "up")
		fmt.Println("Running command:", cmdDBUp.String())
		if err := cmdDBUp.Run(); err != nil {
			fmt.Println("Error running goose up:", err)
			return
		}
	case "dbcreate":
		cmdDBCreate := exec.Command("goose", "-dir", migrationDir, "create", os.Args[2], "sql")
		fmt.Println("Running command:", cmdDBCreate.String())
		cmdDBCreate.Stdout = os.Stdout
		cmdDBCreate.Stderr = os.Stderr
		if err := cmdDBCreate.Run(); err != nil {
			fmt.Println("Error running goose create:", err)
			return
		}
	case "dbversion":
		cmdDBVersion := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "version")
		fmt.Println("Running command:", cmdDBVersion.String())
		cmdDBVersion.Stdout = os.Stdout
		cmdDBVersion.Stderr = os.Stderr
		if err := cmdDBVersion.Run(); err != nil {
			fmt.Println("Error running goose version:", err)
			return
		}
	case "dbstatus":
		cmdDBStatus := exec.Command("goose", "-dir", migrationDir, dbDriver, dbUrl, "status")
		fmt.Println("Running command:", cmdDBStatus.String())
		cmdDBStatus.Stdout = os.Stdout
		cmdDBStatus.Stderr = os.Stderr
		if err := cmdDBStatus.Run(); err != nil {
			fmt.Println("Error running goose version:", err)
			return
		}
	case "dbseed":
		dbSeedCmd := flag.NewFlagSet("dbseed", flag.ExitOnError)
		nUser := dbSeedCmd.Int("nuser", 10, "number of user")
		nOwner := dbSeedCmd.Int("nowner", 10, "number of owner")
		nLocation := dbSeedCmd.Int("nlocation", 10, "number of location")
		nFood := dbSeedCmd.Int("nfood", 10, "number of food")
		dbSeedCmd.Parse(os.Args[2:])

		db.Seed(*nUser, *nOwner, *nLocation, *nFood)
	case "buildlinux":
		buildLinux := exec.Command("GOOS", "=", "linux", "GOARCH", "=", "amd64", "go", "build", "-ldflags", "=", `"-s -s"`, "-o", "jejakmakan-api-linux", "../build")
		fmt.Println("Running command:", buildLinux.String())
		buildLinux.Stdout = os.Stdout
		buildLinux.Stderr = os.Stderr
		if err := buildLinux.Run(); err != nil {
			fmt.Println("Error running build linux:", err)
			return
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
