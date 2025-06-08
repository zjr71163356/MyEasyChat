userprotoc:
	goctl rpc protoc user/proto/user.proto  --go_out=user/rpc  --go-grpc_out=user/rpc --zrpc_out=user/rpc  \
	--proto_path=user/proto 
userapi:
	goctl api go --api user/api/user.api --dir user/api --test

.PHONY: protoc