# Change this to your own Go module
GO_MODULE := my-protobuf

.PHONY: tidy
tidy:
	go mod tidy && go mod vendor  

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif

.PHONY: protoc
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. ./proto/basic/*.proto ./proto/dummy/*.proto ./proto/jobsearch/*.proto

.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run