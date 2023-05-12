# JWT Authenticated API With GO

## Introduction

This is a simple API built with Go and mongo database where you users can sign up as either ADMIN or USER, after successful signup, they get a token. They can then login with their email and password. A user signed up as user will only be able to get his or her details back while the admin can get all the users that signed up or a specific user provided the admin knows the person's user id. This API is protected with JWT authentication.

##### Do the stuff below to initialize your project

1. Clone the repository
   `git clone https://github.com/Shittu24/go-jwt-auth.git`
2. Change the directory
   `cd go-jwt-auth`
3. Get all the dependencies
   `go mod tidy`
4. Run your project
   `go run main.go`

   Then go to postman and start running the routes
