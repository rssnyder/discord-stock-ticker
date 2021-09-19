build:
	env CGO_ENABLED=1 go build -o ./discord-stock-ticker -ldflags="-X 'main.buildVersion=vdev'"

package:
	tar cvfz discord-stock-ticker.tar.gz discord-stock-ticker

run:
	./discord-stock-ticker -logLevel=0

run-dev:
	./discord-stock-ticker -logLevel=1
