build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/ticker

build-osx:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/ticker

run:
	./bin/ticker -logLevel=0

run-dev:
	./bin/ticker -logLevel=1