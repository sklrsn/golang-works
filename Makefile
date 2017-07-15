all: run

tests:
		go test

run:
	 go build *.go
	 go run main.go
