OUT := order_api
VERSION := $(shell git describe --always --long --dirty)
DB := sha

default: start

.PHONY: setup
setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pressly/goose/cmd/goose

.PHONY:
dep:
	@dep ensure -v

.PHONY: build
build:
	mkdir -p build
	go build -o build/${OUT}

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/${OUT}-linux

.PHONY: clean
clean:
	rm -rf build/

.PHONY: start
start:
	build/${OUT}

.PHONY: container
container: clean build-linux
	docker build -t ${OUT}:${VERSION} -f Dockerfile .
	docker tag ${OUT}:${VERSION} ${OUT}:latest

.PHONY: container-start
container-start:
	docker run --env-file ./env.list -d -p 8080:8080 --name ${OUT}.local ${OUT}:${VERSION}

.PHONY: container-stop
container-stop:
	docker rm -f ${OUT}.local

.PHONY: db-container
db-container:
	cd db/ && docker build -t ${DB}:${VERSION} -f Dockerfile .
	docker tag ${DB}:${VERSION} ${DB}:latest

.PHONY: db-container-start
db-container-start:
	docker run -d -p 5432:5432 --name ${DB}.db ${DB}:${VERSION}