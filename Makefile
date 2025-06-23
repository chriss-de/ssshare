#
#
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
NAME=$(CURRENT_DIR)

.PHONY: gomod build run swagger-docs docker-image clean

all: build

gomod:
	go mod tidy
	go mod vendor

build: gomod swagger-docs
	mkdir -p out/bin
	CGO_ENABLED=0 go build -o out/bin/$(NAME) .

run: gomod swagger-docs
	go run .

swagger-docs:
	go run github.com/swaggo/swag/cmd/swag@v1.8.12 init

docker-image:
	buildah bud -f Dockerfile .

clean:
	rm -fr out

