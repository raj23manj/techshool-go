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
