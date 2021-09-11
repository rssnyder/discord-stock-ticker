build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"

build-osx:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"
	
build-openbsd:
	env GOOS=openbsd GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"

run:
	./bin/ticker -logLevel=0

run-dev:
	./bin/ticker -logLevel=1
