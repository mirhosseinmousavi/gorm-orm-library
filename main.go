package main

// Import all packages needed for this project
import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// To start this application please use first go mod tidy and then use to run.
// We must use "go run ." because if you use just "go run main.go"
// The code getting the error because of using struct in another file in the same package
// and when we use go run main.go other files like model.go are ignored and you must consider another
// file is important to you and to use model.go you can use the command below:
// go run main.go model.go
// If you have multiple files this way is not good for you
// and I suggest using " go run . " to contain all go files needed to run when you start the go application
// I hope this code is useful for you
// if you have any questions please contact me
// Use the main function as the entry point of the application
func main() {
	// create database.db if not exist this file and open up this file to use as a database
	// Use the default configuration of gorm we can skip this step right now
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	// If any problem exists we can stop the application with a message like failed to connect and not able to continue
	if err != nil {
		panic("failed to connect")
	}

	// Migrate the model based on the struct we created in model.go file
	// This code will create a table if not exists
	db.AutoMigrate(&CustomerProfile{})

	// Create a new record based on the CustomerProfile model and using CustomerInformation and CustomerScore
	// in another struct to simplify the struct and prevent all model variables in the same struct to manage better
	// other structs
	// using a simple password just for the test
	// Please notice about the hashing password for your project
	// and using a secret key to encrypt and decrypt passwords to generate and verify password
	customerInformation := CustomerInformation{FirstName: "Hossein", LastName: "Mousavi", Email: "mirhosseinmousavi42@gmail.com", Password: "123456"}
	customerScore := CustomerScore{Amount: 10}
	customerProfile := CustomerProfile{CustomerInformation: customerInformation, CustomerScore: customerScore}
	db.Create(&customerProfile)

	// return blog id based on creation id and fill id when we create in previous code
	// We can use this id later to update the same record we created before
	fmt.Printf("customerProfile id: %v\n", customerProfile.ID)

	// as we can use customer_information_ in embeddedPrefix
	// we can join customer_information_ to first_name
	// because we create first_name in another struct
	// and use this in CustomerProfile after that, we can change the value of these fields to other things like
	// in this example use Test as the first name of the customer
	db.Model(&customerProfile).Update("customer_information_first_name", "Test")

	// After we update the first name to test this field is changed we use
	// if record was find with id = 1 we can get this field in CustomerProfile
	// and if not any record was found we take an error about the record not being found
	err = db.First(&customerProfile, 1).Error

	// to check error is a record as not found and not any other error happened
	// We use the code below:
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("can not found any record with id equal to 1")
	}

	// If the record was found we print this record and check everything is changed correctly
	// We use Printf to concatenate the string with the variable and use %v to use the value of the second parameter
	// use instead of %v in when places use of %v
	fmt.Printf("customer profile data after changed first_name: %v\n", customerProfile.CustomerInformation.FirstName)

	// At the end of the use of gorm as orm library, we delete the record we created before to test
	// delete function of this library and test it with the number of deleted records with RowsAffected
	rowsOfDeleted := db.Delete(&customerProfile, 1).RowsAffected

	// and to check the number of deleted records we print rowsOfDeleted
	fmt.Printf("rowsOfDeleted: %v\n", rowsOfDeleted)

}
