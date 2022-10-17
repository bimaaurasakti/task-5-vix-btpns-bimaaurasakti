## Project Layout

```text
├── app
│   │── photo
│   │   ├── formatter.go
│   │   └── input.go
│   │── user
│   │   ├── formatter.go
│   │   └── input.go
│   └── app.go
├── controllers
│   ├── AuthController.go
│   ├── PhotoController.go
│   └── UserController.go
├── database
│   └── database.go
├── controllers
│   ├── api_response.go
│   ├── env.go
│   ├── jwt.go
│   └── validation.go
├── middlewares
│   └── middlewares.go
├── models
│   ├── User.go
│   └── UserPhoto.go
├── public
│   └── images
│       └── user
├── router
│   └── router.go
├── .env.example
├── go.mod
├── go.sum
└── main.go
```
