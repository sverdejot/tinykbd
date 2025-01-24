.DEFAULT_GOAL := all

.PHONY: all run flash build vet fmt build dist

BIN_FILE = "$(PWD)/dist/tinykbd"
INSTALL_SCRIPT = "$(PWD)/install.sh"

all: flash run

flash:
	@tinygo flash -target=xiao firmware/main.go

run: build
	@./tinykbd

build: vet
	@go build -o $(BIN_FILE) ./app/

vet: fmt
	@go vet ./app

fmt:
	@go fmt ./...

dist: build
	@aws s3 cp $(BIN_FILE) s3://$(S3_BUCKET)/$(BIN_FILE_S3_KEY) --region $(AWS_REGION)
	@aws s3 cp $(INSTALL_SCRIPT) s3://$(S3_BUCKET)/$(INSTALL_SCRIPT_S3_KEY) --region $(AWS_REGION)
