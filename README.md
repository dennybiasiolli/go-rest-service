# Go REST Service

#### Setup

- gorilla/mux — A powerful URL router and dispatcher. We use this package to match URL paths with their handlers.

- jinzhu/gorm — The fantastic ORM library for Golang, aims to be developer friendly. We use this ORM(Object relational mapper) package to interact smoothly with our database

- dgrijalva/jwt-go — Used to sign and verify JWT tokens

- joho/godotenv — Used to load .env files into the project

```sh
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
```

#### Start

`go run main.go`

Open http://localhost:8000/products/ or http://localhost:8000/res-users/ in your browser
