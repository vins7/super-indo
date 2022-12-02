export SUPER_INDO_DB_HOST=localhost
export SUPER_INDO_DB_PORT=5306
export SUPER_INDO_DB_USERNAME=root
export SUPER_INDO_DB_PASSWORD=root
export SUPER_INDO_DB_NAME=super_indo

export REDIS_HOST=localhost
export REDIS_PORT=6379

export SUPER_INDO_PORT=8080
export SUPER_INDO_SECRET=super_indo

run:
	@go run main.go
build: 
	@go build -o ./main main.go
	