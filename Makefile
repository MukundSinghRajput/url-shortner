build:
	@echo "Building..."
	@templ generate
	@go build -o bin/us cmd/main.go

run:
	@echo "Starting the App"
	./bin/us

start: build run

test:
	@echo "Testing..."
	@go test ./tests -v