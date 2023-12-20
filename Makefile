.PHONY: run build

air: build
	@go build -o ./tmp/main cmd/main.go

run: build
	@go run cmd/main.go

build:
	@templ generate
