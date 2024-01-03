build:
	@go build -o bin/basic_tuto.go

run: build
	@./bin/basic_tuto

test:
	@go test 