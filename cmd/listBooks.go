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

// listBooksCmd represents the listBooks command
var listBooksCmd = &cobra.Command{
	Use:   "listBooks",
	Aliases: []string{"list", "l", "lb"},
	Short: "List all the books in your library.",
	Long: `Lists all the books in your library.`,
	Run: func(cmd *cobra.Command, args []string) {
		var books []database.Book
		result := database.DB.Find(&books)
		
		if result.Error != nil {
			log.Fatal("Failed to list books", result.Error)
			return
		}

		fmt.Printf("Found %d books:\n\n", result.RowsAffected)
		for index, book := range books {
			fmt.Printf("%d) %s by %s\n", index + 1, book.Title, book.Author)
		}

		fmt.Print("\n")
	},
}

func init() {
	rootCmd.AddCommand(listBooksCmd)
}
