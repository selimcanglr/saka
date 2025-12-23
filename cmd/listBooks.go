/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/selimcanglr/book-cli/internal/database"
	"github.com/spf13/cobra"
)

var listBooksCmd = &cobra.Command{
	Use:   "listBooks",
	Aliases: []string{"list"},
	Short: "List all the books in your library.",
	Long: `You may use this command to see every book in your library.
	
Listed books are ordered by the date they are added.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var books []database.Book
		result := database.DB.
				Order("created_at desc").
				Find(&books)
		
		if result.Error != nil {
			log.Fatal("Failed to list books", result.Error)
			return
		}

		fmt.Printf("Found %d books:\n\n", result.RowsAffected)
		for index, book := range books {
			fmt.Printf("%d) %s by %s\n", index + 1, book.Title, book.Author)
		}
	},
}

func init() {
	rootCmd.AddCommand(listBooksCmd)
}
