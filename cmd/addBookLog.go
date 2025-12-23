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

// TODO: Add sub-command "last" to add a log to the latest book, skipping the book selection step
// TODO: Allow using flags to rate: title, rating, entry?

var addBookLogCmd = &cobra.Command{
	Use:     "addBookLog",
	Aliases: []string{"log"}, // Added handy aliases
	Short:   "Add a reading log/entry for a book",
	Long:    `Record your progress or thoughts while reading a book. Tracks page number and comments.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Fetch Books
		var books []database.Book
		result := database.DB.Find(&books)
		if result.Error != nil {
			log.Fatal("Failed to fetch books")
		}

		if len(books) == 0 {
			fmt.Println("No books found. Add a book first!")
			return
		}

		// 2. Prepare Options
		var bookOptions []huh.Option[uint]
		for _, b := range books {
			bookOptions = append(bookOptions, huh.NewOption(b.Title, b.ID))
		}

		// 3. Define Form Variables
		var selectedBookID uint
		var pageStr string
		var entryText string

		// 4. Create the Form
		form := huh.NewForm(
			// Group 1: Select Book
			huh.NewGroup(
				huh.NewSelect[uint]().
					Title("Which book are you reading?").
					Options(bookOptions...).
					Value(&selectedBookID),
			),

			// Group 2: Log Details
			huh.NewGroup(
				// Input for Page Number (with validation)
				huh.NewInput().
					Title("Current Page").
					Placeholder("e.g. 42").
					Value(&pageStr).
					Validate(func(str string) error {
						if _, err := strconv.Atoi(str); err != nil {
							return fmt.Errorf("please enter a valid number")
						}
						return nil
					}),

				// Text area for the log entry
				huh.NewText().
					Title("Log Entry").
					Placeholder("Any thoughts?").
					CharLimit(3000).
					Value(&entryText),
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
		pageInt, _ := strconv.Atoi(pageStr) // Error ignored because we validated it above

		newLog := database.BookLog{
			BookID: selectedBookID,
			Page:   uint(pageInt),
			Entry:  entryText,
		}

		if err := database.DB.Create(&newLog).Error; err != nil {
			log.Fatal("Failed to save log:", err)
		}

		fmt.Printf("\nLogged page %d for book ID %d.\n", pageInt, selectedBookID)
	},
}

func init() {
	rootCmd.AddCommand(addBookLogCmd)
}