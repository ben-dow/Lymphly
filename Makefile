.PHONY: build 

APPLICATION_NAME=Lymphly

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_DIR=${ROOT_DIR}/.build
RELEASE_DIR=${BUILD_DIR}/releases

build_dir:
	mkdir -p ${BUILD_DIR} ${RELEASE_DIR}

build/lambda/providersearch: build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${BUILD_DIR}/providersearch/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/providersearch/main.go
	cd ${BUILD_DIR}/providersearch; \
	zip providersearch_lambda_x86_64.zip bootstrap; \
	mv providersearch_lambda_x86_64.zip ${RELEASE_DIR}/providersearch_lambda_x86_64.zip


build/lambda/providerupdate: build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${BUILD_DIR}/providerupdate/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/providerupdate/main.go
	cd ${BUILD_DIR}/providerupdate; \
	zip providerupdate_lambda_x86_64.zip bootstrap; \
	mv providerupdate_lambda_x86_64.zip ${RELEASE_DIR}/providerupdate_lambda_x86_64.zip

build/frontend: build_dir
	cd frontend; \
	npm run build; \
	zip -r dist.zip dist/; \
	mv dist.zip ${RELEASE_DIR}/frontend.zip

build: build/frontend build/lambda/providersearch build/lambda/providerupdate

clean:
	rm -rf .build