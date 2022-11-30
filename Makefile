# generate all protos
gnpall:
	protoc --go_out=grpc-app/protos/auth --go_opt=Mprotos/auth.proto=/ --go-grpc_out=grpc-app/protos/auth --go-grpc_opt=Mprotos/auth.proto=/ protos/auth.proto
	protoc --go_out=grpc-app/protos/users --go_opt=Mprotos/user.proto=/ --go-grpc_out=grpc-app/protos/users --go-grpc_opt=Mprotos/user.proto=/ protos/user.proto
	protoc --go_out=grpc-app/protos/products --go_opt=Mprotos/product.proto=/ --go-grpc_out=grpc-app/protos/products --go-grpc_opt=Mprotos/product.proto=/ protos/product.proto
	protoc --proto_path=protos protos/image.proto --go_out=grpc-app/protos/images --go-grpc_out=grpc-app/protos/images

# generate all injector
injall:
	wire ./grpc-app/injector/users
	wire ./grpc-app/injector/auth
	wire ./grpc-app/injector/products

# genrate all mock service
mock:
	mockgen -package=mock_repo -source=shared/repository/products/product.go -destination=shared/repository/products/mocks/product_repo_mock.go
	mockgen -package=mock_repo -source=shared/repository/users/user.go -destination=shared/repository/users/mocks/user_repo_mock.go
	mockgen -package=mock_repo -source=shared/repository/auth/auth.go -destination=shared/repository/auth/mocks/auth_repo_mock.go
	mockgen -package=mock_repo -source=shared/usecase/products/product.go -destination=shared/usecase/products/mocks/product_usecase_mock.go
	mockgen -package=mock_repo -source=shared/usecase/users/user.go -destination=shared/usecase/users/mocks/user_usecase_mock.go
	mockgen -package=mock_repo -source=shared/usecase/auth/auth.go -destination=shared/usecase/auth/mocks/auth_usecase_mock.go
	mockgen -package=mock_repo -source=grpc-app/service/auth/auth.go -destination=grpc-app/service/auth/mocks/auth_service_mock.go
	mockgen -package=mock_repo -source=grpc-app/service/products/product.go -destination=grpc-app/service/products/mocks/product_service_mock.go
	mockgen -package=mock_repo -source=grpc-app/service/users/user.go -destination=grpc-app/service/users/mocks/user_service_mock.go

# docker build image
bimage:
	docker build -t server-alpha .

# build container
bcon:
	docker container create --name server-alpha --net grpc server-alpha

#run docker
rdoc:
	docker start server-alpha

# run migration
migration:
	go run ./cmd/postgres/main.go

# run all test
test:
	go test ./... -coverprofile coverage.out
	rm -rf coverage.out

# run go
r:
	go run main.go

# runing all test before run app
run: test r run