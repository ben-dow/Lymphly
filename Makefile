.PHONY: build 

APPLICATION_NAME=Lymphly

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_DIR=${ROOT_DIR}/.build
RELEASE_DIR=${BUILD_DIR}/releases

build_dir:
	mkdir -p ${BUILD_DIR} ${RELEASE_DIR}

build/lambda/public: build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${BUILD_DIR}/public/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/public/main.go
	cd ${BUILD_DIR}/public; \
	zip public_lambda_x86_64.zip bootstrap; \
	mv public_lambda_x86_64.zip ${RELEASE_DIR}/public_lambda_x86_64.zip

build/lambda/private: build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ${BUILD_DIR}/private/bootstrap -ldflags "-s -w" ${ROOT_DIR}/cmd/private/main.go
	cd ${BUILD_DIR}/private; \
	zip private_lambda_x86_64.zip bootstrap; \
	mv private_lambda_x86_64.zip ${RELEASE_DIR}/private_lambda_x86_64.zip

build/frontend: build_dir
	cd frontend; \
	npm install; \
	npm run build; \
	zip -r dist.zip dist/; \
	mv dist.zip ${RELEASE_DIR}/frontend.zip

build: build/frontend build/lambda/public build/lambda/private

clean:
	rm -rf .build