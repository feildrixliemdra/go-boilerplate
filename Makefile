install:
	go mod download && go mod vendor


dev:
	go run main.go serve-http