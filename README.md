# transporte-backend


## Requirements

- PostgresSQL >= v10
- Golang >= 1.13


## Dependencies

- [Echo Framework](https://echo.labstack.com/)
- [Gorm ORM](http://gorm.io/es_ES/docs/index.html)

### install dependencies

For manager dependencies used `go mod`, for download run in your bash console:
```bash
$ go mod download
```

## Set database params

edit connection params into  database/connection.go file in `init` method.

## migrations
When the server stars run automatically all migrations.

## Settings

### admin user
If admin user not exist, automatically register a new user
```text
username: admin
password: admin
``` 
is recommended change the password in the first login.

## run

```bash
$ go run .
```

## Authors

- Fabio Moreno <FabioMoreno@outlook.com>