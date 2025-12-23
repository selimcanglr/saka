/*
Copyright © 2025 Selim Can Güler <cs.selim.guler@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/selimcanglr/book-cli/internal/database"
	"github.com/spf13/cobra"
)

var title string
var author string

// TODO: Display huh form if no flag is provided

var addBookCmd = &cobra.Command{
	Use:   "addBook",
	Aliases: []string{"add"},
	Short: "Add a new book",
	Long: `Adds a new book to your library.`,
	Run: func(cmd *cobra.Command, args []string) {
		newBook := database.Book{
			Title: title,
			Author: author,
		}

		result := database.DB.Create(&newBook)
		if result.Error != nil {
			log.Fatal("Failed to add to DB: ", result.Error)
		}

		fmt.Printf("Added book successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addBookCmd)

	addBookCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the book")
	addBookCmd.MarkFlagRequired("title")

	addBookCmd.Flags().StringVarP(&author, "author", "a", "", "Author of the book")
	addBookCmd.MarkFlagRequired("author")
}
