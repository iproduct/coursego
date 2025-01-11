# 1. Generate protocol buffers go bindings for todo_service.proto
mkdir -p generated
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=./generated --go-grpc_out=./generated todo_service.proto


# 2. Add swagger API
go get -u github.com/golang/protobuf/protoc-gen-go
mkdir -p api\swagger\v1

protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=./generated --go-grpc_out=./generated todo_service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:generated todo_service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 todo_service.proto

## All in one generation
D:\CourseGO\protoc-3.14.0-win64\bin\protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=./generated --go-grpc_out=./generated --grpc-gateway_out=logtostderr=true:generated --swagger_out=logtostderr=true:api/swagger/v1 todo_service.proto