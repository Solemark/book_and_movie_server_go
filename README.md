# Book and Movie Server - Go

run code with:

    go run books/main.go
    go run movies/main.go
    go run reverse_proxy/main.go
    go run client/main.go

build with:
	
	go build name.go
	OR
	go build -ldflags -w name.go

-ldflags -w removes debug symbols resulting in a smaller binary

target another system with:

	GOOS=platform GOARCH=system go build name.go
	e.g
	GOOS=windows GOARCH=amd64 go build name.go
