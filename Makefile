start:
	@clear
	@echo ▀█▀ █ █▀█ ▀▄▀
	@echo
	@go build -o bin/goblog cmd/main.go
	@./bin/goblog

tidy:
	@go mod tidy

test:
	@go test -v ./...