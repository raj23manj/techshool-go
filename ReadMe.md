# Make file
  * add all the commands to be used by team members (https://github.com/techschool/simplebank)
  * use command to run easily `make migrateup`
# Migration
  * `migrate create -ext sql -dir db/migration -seq init_schema`
  * `migrate -path db/migration -database "postgres://rajeshmanjunath:password@localhost:5432/simple_bank?sslmode=disable" -verbose up`
# DB pkg golang
  * https://pkg.go.dev/database/sql#DB.QueryContext
# ORM
  * SQLC
  ## SQLC commands
    * sqlc init
# Initialise the project to avoid errors
  * go mod init
  * go mod tidy => to install dependencies
# Packages used in this app
  * db driver package to connect to database while testing
    - github.com/lib/pq
  * Test assertion & mocking with unit testing
    - github.com/stretchr/testify
  * app builds own fake data generation, but see
    - github.com/brianvoe/gofakeit
    - https://medium.com/the-bug-shots/generating-gofakeit-data-for-testing-in-go-with-gofakeit-2726b5fd4fa9
  * Testing Tear down & setup db
    - https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
