user_protoc:
	goctl rpc protoc user/proto/user.proto  --go_out=user/rpc  --go-grpc_out=user/rpc --zrpc_out=user/rpc  \
	--proto_path=user/proto 
user_api:
	goctl api go --api user/api/user.api --dir user/api --test
user_models:
	goctl model  mysql ddl -s /home/tyrfly/MyEasyChat/user/sql/user.sql -dir /home/tyrfly/MyEasyChat/user/models -c
	
run_mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=azsx0123456  -p 3307:3306 -d mysql:latest
.PHONY: protoc models api