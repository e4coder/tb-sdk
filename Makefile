.PHONY: check-gocritic run-gocritic check-gosec run-gosec run-tests check-golangci-lint run-golangci-lint

all: check-gocritic run-gocritic check-gosec run-gosec check-golangci-lint run-golangci-lint run-tests 

check-gocritic:
	@command -v gocritic >/dev/null 2>&1 || { echo >&2 "gocritic is not installed. Please install it."; exit 1; }

run-gocritic:
	gocritic check ./...

check-gosec:
	@command -v gosec >/dev/null 2>&1 || { echo >&2 "gosec is not installed. Please install it."; exit 1; }

run-gosec:
	gosec ./...

check-golangci-lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo >&2 "golangci-lint is not installed. Please install it."; exit 1; }

run-golangci-lint:
	golangci-lint run ./...

run-tests:
	@echo "Checking health of bundler..."
	@HEALTH_CHECK=$$(curl -s "localhost:3000/health"); \
	if [ "$$HEALTH_CHECK" = "\"ok\"" ]; then \
	    go test -v ./...; \
	else \
	    echo "Bundler service is not running."; \
	    exit 1; \
	fi

get-coverage:
	go test -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html
	open coverage.html
