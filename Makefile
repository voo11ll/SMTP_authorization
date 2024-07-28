build-all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/user_http cmd/http/user/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/user_grpc cmd/grpc/user/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/organization_http cmd/http/organization/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/organization_grpc cmd/grpc/organization/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/notification_grpc cmd/grpc/notification/main.go
	docker-compose -f "docker-compose.yaml" up -d --build

restart-docker:
	docker-compose -f "docker-compose.yaml" up -d --build

build-user:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/user_http cmd/http/user/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/user_grpc cmd/grpc/user/main.go
	docker-compose -f "docker-compose.yaml" up -d --build b24-user-grpc
	docker-compose -f "docker-compose.yaml" up -d --build b24-user-http
	docker restart b24-user-http
	docker restart b24-user-grpc

build-notification:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/notification_grpc cmd/grpc/notification/main.go
	docker-compose -f "docker-compose.yaml" up -d --build b24-notification-grpc
	docker restart b24-user-notification

build-organization:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/organization_http cmd/http/organization/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/organization_grpc cmd/grpc/organization/main.go
	docker restart organization-http
	docker restart organization-grpc

.DEFAULT_GOAL := build-all