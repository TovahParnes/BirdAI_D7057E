up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker build -t bird_ai .

con:
	docker container ls

run:
	swag init -g main.go --output src/internal/docs
	go run main.go
