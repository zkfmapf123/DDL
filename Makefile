build:
	@rm -rf ddl
	@echo "Building for x86 architecture"
	@CGO_ENABLED=0 GOARCH=amd64 go build -o bin/x86/ddl
	@echo "Building for ARM architecture"
	@CGO_ENABLED=0 GOARCH=arm64 go build -o bin/arm/ddl

run:
	go run main.go

which:
	mv ./bin/arm/ddl .

start: build
	@make which
	./ddl

## Install golnagci-lint
lint:
	golangci-lint run ./...

test: lint
	go test -v ./internal
