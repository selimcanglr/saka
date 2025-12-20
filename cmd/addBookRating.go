/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/selimcanglr/book-cli/internal/database"
	"github.com/spf13/cobra"
)

// addBookRatingCmd represents the addBookRating command
var addBookRatingCmd = &cobra.Command{
	Use:   "addBookRating",
	Aliases: []string{"rate", "ar", "abr"},
	Short: "Add a rating to an existing book",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Fetch data
		var books []database.Book
		result := database.DB.Find(&books)
		if result.Error != nil {
			log.Fatal("Failed to fetch books")
		}

		if len(books) == 0 {
			fmt.Println("No books found in the database.")
			return
		}

		// Prepare Options for huh
		var bookOptions []huh.Option[uint]
		for _, b := range books {
			label := fmt.Sprintf("%s (by %s)", b.Title, b.Author)
			bookOptions = append(bookOptions, huh.NewOption(label, b.ID))
		}

		
		// Create and Run the Form
		var selectedBookID uint
		var selectedRating string
		form := huh.NewForm(
			// Group 1: Select the Book
			huh.NewGroup(
				huh.NewSelect[uint]().
					Title("Which book would you like to rate?").
					Options(bookOptions...).
					Value(&selectedBookID),
			),

			// Group 2: Pick the Rating
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("How was it?").
					Options(
						huh.NewOption("★★★★★ (5/5) - Amazing", "5"),
						huh.NewOption("★★★★☆ (4/5) - Good", "4"),
						huh.NewOption("★★★☆☆ (3/5) - Okay", "3"),
						huh.NewOption("★★☆☆☆ (2/5) - Bad", "2"),
						huh.NewOption("★☆☆☆☆ (1/5) - Terrible", "1"),
					).
					Value(&selectedRating),
			),
		)

		err := form.Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				os.Exit(0)
			}
			log.Fatal(err)
		}

		// 5. Save to Database
		ratingInt, _ := strconv.Atoi(selectedRating)

		newRating := database.BookRating{
			BookID: selectedBookID,
			Rating: uint8(ratingInt),
		}

		insertRatingResult := database.DB.Create(&newRating)
		if insertRatingResult.Error != nil {
			log.Fatal("Failed to save rating:", insertRatingResult.Error)
		}

		book, found := findBookById(books, selectedBookID)
		if !found {
			log.Fatal("Selected book does not exist")
		}

		fmt.Printf("\nSaved %d star rating for book %s.\n", ratingInt, book.Title)
	},
}

func findBookById(books []database.Book, id uint) (database.Book, bool) {
    for _, book := range books {
        if book.ID == id {
            return book, true // Found it
        }
    }

    return database.Book{}, false 
}

func init() {
	rootCmd.AddCommand(addBookRatingCmd)
}