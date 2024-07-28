.PHONY: proto
gen-all:
	protoc --go_out=../pkg/services/grpc/user/ --go_opt=paths=source_relative --go-grpc_out=../pkg/services/grpc/user/  --go-grpc_opt=paths=source_relative user.proto
	protoc --go_out=../pkg/services/grpc/organization/ --go_opt=paths=source_relative --go-grpc_out=../pkg/services/grpc/organization/  --go-grpc_opt=paths=source_relative organization.proto
	protoc --go_out=../pkg/services/grpc/notification/ --go_opt=paths=source_relative --go-grpc_out=../pkg/services/grpc/notification/  --go-grpc_opt=paths=source_relative notification.proto
.DEFAULT_GOAL := gen-all