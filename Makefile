.PHONY: run build
run: build
	@go run cmd/main.go

build:
	@templ generate
