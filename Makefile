userprotoc:
	goctl rpc protoc user/proto/user.proto  --go_out=user/rpc  --go-grpc_out=user/rpc --zrpc_out=user/rpc  \
	--proto_path=user/proto 
userapi:
	goctl api go --api user/api/user.api --dir user/api --test
usermodels:
	goctl model  mysql ddl -s /home/tyrfly/MyEasyChat/user/sql/user.sql -dir /home/tyrfly/MyEasyChat/user/models -c

.PHONY: protoc models api