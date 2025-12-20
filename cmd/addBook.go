/*
Copyright © 2025 Selim Can Güler <cs.selim.guler@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/selimcanglr/book-cli/internal/database"
	"github.com/spf13/cobra"
)

var addBookCmd = &cobra.Command{
	Use:   "addBook",
	Aliases: []string{"add", "ab"},
	Short: "Add a new book",
	Long: `Adds a new book to your library.`,
	Run: func(cmd *cobra.Command, args []string) {
		var title string
		var author string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Title:").
					Description("Book title").
					Placeholder("Software Engineering at Google").
					Value(&title),
				huh.NewInput().
					Title("Author:").
					Description("Author of the book").
					Placeholder("Titus Winters").
					Value(&author),
			),
		)

		err := form.Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				os.Exit(0)
			}
			log.Fatal(err)
		}

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
	Args: cobra.MatchAll(cobra.OnlyValidArgs),
}

func init() {
	rootCmd.AddCommand(addBookCmd)
}
