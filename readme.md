Following executable needs to be installed:

MacOS: brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
Info
All protobuf (version 3) definitions will be located in this folder and will eventually be moved to a centralized repository.

Generated protobuf code will be located in gen/pb/*.

Generated OpenAPI v3 documentation will be located in gen/apidoc/{serviceName}/*.