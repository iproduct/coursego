protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=./generated --go-grpc_out=./generated todo_service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:generated todo_service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 todo_service.proto