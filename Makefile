run:
	go run main.go

install:
	dep ensure

run:
	docker-compose up -d

stop:
	docker-compose down
