build-linux:
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"

build-osx:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"
	
build-openbsd:
	env CGO_ENABLED=1 GOOS=openbsd GOARCH=amd64 go build -o ./bin/ticker -ldflags="-X 'main.buildVersion=vdev'"

run:
	./bin/ticker -logLevel=0

run-dev:
	./bin/ticker -logLevel=1
