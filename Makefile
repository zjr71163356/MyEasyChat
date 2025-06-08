userprotoc:
	goctl rpc protoc user/proto/user.proto  --go_out=user/rpc  --go-grpc_out=user/rpc --zrpc_out=user/rpc  \
	--proto_path=user/proto 
.PHONY: protoc