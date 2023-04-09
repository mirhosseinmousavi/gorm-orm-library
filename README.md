# Using gorm as an orm library in Golang 👋

In this repository, we go through using orm library like gorm to handle crud operation in the database

To start this application please use first `go mod tidy` and then use `go run .` to run.

We must use `go run .` because if you use just `go run main.go`

The code getting the error because of using struct in another file in the same package

and when we use `go run main.go` other files like model.go are ignored and you must consider another

the file is important to you and to use model.go you can use the command below:

`go run main.go model.go`

If you have multiple files this way is not good for you

and I suggest you using `go run .` to contain all go files needed to run when you start the go application

I hope this code is useful for you

if you have any questions please contact me

[EmailAddress](mailto:mirhosseinmousavi42@gmail.com)

[MyWebSite](https://orser.ir)
