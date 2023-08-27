# Due to AWS go 1.x depreciation on Dec 31 2023, now the go lambda functions can be deployed only into a
# Custom Lambda Runtime.
# Therefore, the name of the executable needs to be 'bootstrap' and the RPC lib can be excluded as no more needed (ie: '-tags lambda.norpc').
BIN_DIR ?= bin
BINARY_NAME ?= bootstrap

AWS_LAMBDA_FUNC_NAME ?= stop_instances_v2
AWS_LAMBDA_FUNC_ROLE_ARN ?= ''

clean:
	if [ -d ${BIN_DIR} ]; then rm -rf ${BIN_DIR}; fi
	mkdir -p ${BIN_DIR}

build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ${BIN_DIR}/${BINARY_NAME} main.go

zip: build
	cd ${BIN_DIR} && zip ${BINARY_NAME}.zip ${BINARY_NAME}

aws-create:
	aws lambda create-function --function-name ${AWS_LAMBDA_FUNC_NAME} \
		--runtime provided.al2 --handler ${BINARY_NAME} \
		--architectures x86_64 \
		--role ${AWS_LAMBDA_FUNC_ROLE_ARN} \
		--zip-file fileb://${BIN_DIR}/${BINARY_NAME}.zip

aws-deploy-new-code:
	aws lambda update-function-code \
		--function-name ${AWS_LAMBDA_FUNC_NAME} \
		--zip-file fileb://${BIN_DIR}/${BINARY_NAME}.zip