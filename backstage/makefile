BINARY_NAME=main.out

.PHONY: help
## help: show all comment
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## runbuild: run application
runbuild:
	${BINARY_NAME} run

## build: build the application
build:cmd/app/main.go #dependencies 在執行指令之前會進行檢查
	go build -o ${BINARY_NAME} cmd/app/main.go

## run: run go dev
run:
	go run cmd/app/main.go run

## migrateup: migrate up database 
migrateup: 
	migrate -database 'sqlite3://./sqlite-database.db'   -path ./database/migrations force 1
	
## mock: mock all
mock:
	mockgen -source=./domain/leave.go -destination=./internal/leave/repository/mocks/leave_repository.go -package=mocks
	mockgen -source=./domain/user.go -destination=./internal/identification/repository/mocks/user_repository.go -package=mocks

## clear: go clean and remove build file
clean:
	go clean
	rm ${BINARY_NAME}