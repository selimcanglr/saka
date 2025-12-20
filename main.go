package main

import (
	"github.com/selimcanglr/book-cli/cmd"
	"github.com/selimcanglr/book-cli/internal/database"
)

func main() {
	database.InitDB()

	cmd.Execute()
}
