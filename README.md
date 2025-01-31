# Book and Movie Server - Go

run code with:

    cd books && go run .
    cd movies && go run .
    cd coordinator && go run .
    cd client && go run .

run tests with:

    go test
    OR
    go test -v

build with:
	
	go build name.go
	OR
	go build -ldflags -w name.go

-ldflags -w removes debug symbols resulting in a smaller binary

target another system with:

	GOOS=platform GOARCH=system go build name.go
	e.g
	GOOS=windows GOARCH=amd64 go build name.go
