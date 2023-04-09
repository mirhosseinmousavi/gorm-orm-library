package main

import "gorm.io/gorm"

// create CustomerInformation structure to use this as embedded fields in the CustomerProfile struct
type CustomerInformation struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Create another structure as fields of CustomerProfile to manage the score of a user in CustomerProfile fields
type CustomerScore struct {
	Amount int
}

// Write the main struct of the customer profile to handle the detail of all things related to the customer profile to
// Manage and handle all information needed about every customer
type CustomerProfile struct {
	// You notice that when we use gorm.Model some information need to add by myself added automatically
	// and don't need to type these fields by myself in this struct we have some important fields like
	// ID, CreatedAt, UpdatedAt, DeletedAt and we can use these fields just by gorm.Model
	gorm.Model
	CustomerInformation CustomerInformation `gorm:"embedded;embeddedPrefix:customer_information_"`
	CustomerScore       CustomerScore       `gorm:"embedded;embeddedPrefix:customer_score_"`
}

// We use this function and assign it to CustomerProfile as a method to change the table name based on my prefer table name
func (CustomerProfile) TableName() string {
	return "profile_table"
}
