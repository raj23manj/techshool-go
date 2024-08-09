# Make file
  * add all the commands to be used by team members (https://github.com/techschool/simplebank)
  * use command to run easily `make migrateup`
# Migration
  * `migrate create -ext sql -dir db/migration -seq <name>`
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
# Testing and assertion frameworks:
  * gocheck
  * testify
  * gunit
# Packages used in this app
  * db driver package to connect to database while testing
    - github.com/lib/pq => old use pgx
  * Test assertion & mocking with unit testing
    - github.com/stretchr/testify
  * app builds own fake data generation, but see
    - github.com/brianvoe/gofakeit
    - https://medium.com/the-bug-shots/generating-gofakeit-data-for-testing-in-go-with-gofakeit-2726b5fd4fa9
  * Testing Tear down & setup db
    - https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
    - https://www.gopherguides.com/articles/test-cleanup-in-go-1-14
    - https://stackoverflow.com/questions/42310088/setup-and-teardown-for-each-test-using-std-testing-package
    - https://stackoverflow.com/questions/61609085/what-is-useful-for-t-cleanup
    - https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go/65428147#65428147
    - https://medium.com/kare-nuorteva/go-unit-test-setup-and-teardown-db1601a796f2#.2aherx2z5
    - https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd

# Gin SQLX example:
  - https://github.com/pengenpaham/gin-sqlx-example
# Debugging with context name:
  - see section 1, 9
# DB deadlocks:
  * if two accounts transactions are happening to and from and vice versa, it is always safe to excute the update statements in order
  * update account1 and account 2, so that the locks are taken on first sql and the other goroutine waits until it is done
  * see section, 10
# Transaction levels in PSQL/MYSQL(11)
  * Read Uncommitted
  * Read Committed
  * Repeatable Read
  * Serializable (max)

# Popular web frameworks
  * Gin, Beego, Echo, Revel, Martini, Fiber, Buffalo
  * Popular HTTP routers
    - FastHttp
    - Gorilla Mux
    - HttpRouter
    - chi
# Mock DB for testing
  * Use Fake DB: Memory, implement a fake version of db & store data in memory, like map/hashmap
  * used stubbing, GOMOCK
  * using mockgen:
    - after installing, vi ~/.zshrc and add` export PATH=$PATH:~/go/bin` and `source ~/.zshrc`
   - `mockgen -package mockDB -destination db/mock/store.go  github.com/raj23manj/techshool-go/db/sqlc Store`

# Gvm
  * https://dev.to/0xankit/go-version-manager-gvm-41e7
