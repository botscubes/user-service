all: run


run: 
	go run cmd/server.go 



composeup: docker-compose.yml
	docker-compose up -d

composedown: 
	docker-compose down

gotest: 
	go test ./test