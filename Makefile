up:
	docker-compose up -d

down:
	docker-compose down

con:
	docker container ls

run:
	swag init -g main.go --output src/internal/docs
	go run main.go
