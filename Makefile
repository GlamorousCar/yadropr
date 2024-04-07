
build:
	go build -o dist/xkcd cmd/xkcd/main.go

run:
	$ go run cmd/xkcd/main.go

run-with-o:
	$ go run cmd/xkcd/main.go -o

run-with-n:
	$ go run cmd/xkcd/main.go -n 15

run_xkcd:
	$ .dist/xkcd

