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
	protoc --go_out=. ./proto/basic/*.proto ./proto/dummy/*.proto ./proto/jobsearch/*.proto
ifeq ($(OS), Windows_NT)
	xcopy .\my-protobuf . /E
	if exist "my-protobuf" rd /s /q my-protobuf	
else
	cp -fR ./my-protobuf .
	rm -fR ./my-protobuf	
endif

.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run