.PHONY: tidy
tidy:
	go mod tidy && go mod vendor  

# Windows
.PHONY: clean
clean:
	if exist "protogen" rd /s /q protogen	

# Linux / Mac
# .PHONY: clean
# clean:
#	rm -fR ./protogen

# Windows
.PHONY: protoc
protoc:
	protoc --go_out=. ./proto/basic/*.proto ./proto/dummy/*.proto ./proto/jobsearch/*.proto
	xcopy .\my-protobuf . /E
	if exist "my-protobuf" rd /s /q my-protobuf	

# Linux / Mac
# .PHONY: protoc
# protoc:
#	protoc --go_out=. ./proto/basic/*.proto ./proto/dummy/*.proto ./proto/jobsearch/*.proto
#	cp -fR ./my-protobuf .
#	rm -fR ./my-protobuf	

.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run