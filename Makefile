.PHONY: build 

APPLICATION_NAME=ProviderSearch

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_DIR=${ROOT_DIR}/.build
RELEASE_DIR=${BUILD_DIR}/releases

prep:
	mkdir -p ${BUILD_DIR} ${RELEASE_DIR}

build/lambda:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${RELEASE_DIR}/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/providersearch/main.go
	cd ${RELEASE_DIR}; \
	zip ${APPLICATION_NAME}_lambda_x86_64.zip bootstrap; \
	rm bootstrap

build: prep build/lambda

clean:
	rm -rf .build