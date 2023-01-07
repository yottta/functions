BIN_DIR ?= "bin/"

clean:
	if [ -e ${BIN_DIR} ]; then rm -rf ${BIN_DIR}; fi
	mkdir -p ${BIN_DIR}

build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_DIR}/app main.go

zip: build
	cd ${BIN_DIR} && zip app.zip app
