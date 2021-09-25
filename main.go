package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args

	if envLoadErr := godotenv.Load(".env"); envLoadErr != nil {
		panic("Unable to load .env file make sure it is exists.")
	}
	if len(args) > 1 {
		switch args[1] {
		case "serve":
			StartServer()
		case "db:migrate":
			MigrateDatabase()
		default:
			PrintCommandUsage()
		}
	} else {
		PrintCommandUsage()
	}

}

func PrintCommandUsage() {
	fmt.Print(`
./main [commands]

Available commands
====
serve - Start Server
db:migrate - Start DB Migration

`)

}
