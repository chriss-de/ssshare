#
#
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
NAME=$(CURRENT_DIR)

.PHONY: build run swagger-docs docker-image vendor clean

all: build

build: vendor embedded swagger-docs
	mkdir -p out/bin
	CGO_ENABLED=0 go build -o out/bin/$(NAME) .

run: vendor embedded swagger-docs
	go run .

swagger-docs:
	go run github.com/swaggo/swag/cmd/swag@v1.8.12 init

docker-image:
	buildah bud -f Dockerfile .

vendor:
	go mod vendor

clean:
	rm -fr out

