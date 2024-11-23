_build:
	@rm -rf ddl
	@echo "Building for x86 architecture"
	@CGO_ENABLED=0 GOARCH=amd64 go build -o bin/x86/ddl
	@echo "Building for ARM architecture"
	@CGO_ENABLED=0 GOARCH=arm64 go build -o bin/arm/ddl

dev_build:
	go build -o ddl main.go

run:
	go run main.go

which:
	mv ./bin/arm/ddl .

start: _build
	@make which
	./ddl

dev_start: dev_build
	./ddl

## Install golnagci-lint
lint:
	golangci-lint run ./...

test: lint
	go test -v ./internal
