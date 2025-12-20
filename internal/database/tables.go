package database

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title string
	Author string

	// one to one relationship with BookRating
	Rating BookRating
}

type BookRating struct {
	gorm.Model
	BookID uint

	Rating uint8
	Review string
}

type BookLog struct {
	gorm.Model
	BookID uint

	Entry string
	Page uint
}
