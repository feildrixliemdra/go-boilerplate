install:
	go mod download && go mod vendor


dev:
	go run main.go serve-http


env:
	cp config/config.yaml.example config/config.yaml
	