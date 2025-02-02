.PHONY: build 

APPLICATION_NAME=Lymphly

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_DIR=${ROOT_DIR}/.build
RELEASE_DIR=${BUILD_DIR}/releases

prep:
	mkdir -p ${BUILD_DIR} ${RELEASE_DIR}

build/lambda/providersearch:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${RELEASE_DIR}/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/providersearch/main.go
	cd ${RELEASE_DIR}; \
	zip providersearch_lambda_x86_64.zip bootstrap; \
	rm bootstrap

build/lambda/providerupdate:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${RELEASE_DIR}/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/providerupdate/main.go
	cd ${RELEASE_DIR}; \
	zip providerupdate_lambda_x86_64.zip bootstrap; \
	rm bootstrap

build: prep build/lambda/providersearch build/lambda/providerupdate

clean:
	rm -rf .build