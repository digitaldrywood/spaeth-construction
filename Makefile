SHELL := /bin/bash

.PHONY: dev build test lint generate css css-watch setup clean run help

BINARY_NAME=spaeth-construction

dev:
	@if [ -f tmp/air-combined.log ]; then \
		mv tmp/air-combined.log tmp/air-combined-$$(date +%Y%m%d-%H%M%S).log; \
	fi
	@ls -t tmp/air-combined-*.log 2>/dev/null | tail -n +6 | xargs rm -f 2>/dev/null || true
	@air 2>&1 | tee tmp/air-combined.log

build: generate css
	go build -o $(BINARY_NAME) ./cmd/server

test:
	go test -v -race ./...

lint:
	golangci-lint run
	templ fmt templates/

generate:
	go generate ./...

css:
	npx @tailwindcss/cli -i public/static/css/input.css -o public/static/css/output.css --minify

css-watch:
	npx @tailwindcss/cli -i public/static/css/input.css -o public/static/css/output.css --watch

setup:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

clean:
	rm -f $(BINARY_NAME)
	rm -rf tmp/
	rm -f public/static/css/output.css

run: build
	./$(BINARY_NAME)

help:
	@echo "Available targets:"
	@echo "  dev            - Run with Air hot reload"
	@echo "  build          - Build the binary"
	@echo "  test           - Run tests"
	@echo "  lint           - Run golangci-lint and templ fmt"
	@echo "  generate       - Generate templ code"
	@echo "  css            - Build Tailwind CSS"
	@echo "  css-watch      - Watch and rebuild Tailwind CSS"
	@echo "  setup          - Install development tools"
	@echo "  clean          - Remove build artifacts"
	@echo "  run            - Build and run the server"
