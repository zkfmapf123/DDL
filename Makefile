build: 
    CGO_ENABLED=1 GOARCH=amd64 go build -o bin/x86/ddl
    CGO_ENABLED=1 GOARCH=amd64 go build -o bin/arm/ddl

run:
	go run main.go

which:
	mv /bin/arm/ddl /usr/loca/bin/ddl

start: clean
	make build
	make which
	ddl

test:
	go test -v ./...