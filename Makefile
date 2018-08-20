
default: start

.PHONY: setup
setup:
	go get -u github.com/pressly/goose/cmd/goose

.PHONY: build
build:
	mkdir -p build
	go build -o build/lalamove

.PHONY: clean
clean:
	rm -rf build/

.PHONY: start
start:
	build/lalamove