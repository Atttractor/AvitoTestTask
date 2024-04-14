server:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go avito-app

d.build:
	docker build -t avito-app

d.up.build:
	docker-compose up --build avito-app

d.down:
	docker-compose down

d.up:
	docker-compose up