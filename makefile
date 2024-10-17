generate-post-service-api:
	mkdir -p services/post-service/proto/post-service
	protoc -I services/post-service/api/post-service -I vendor-protogen \
      --go_out=services/post-service/proto/post-service --go_opt=paths=source_relative \
      --go-grpc_out=services/post-service/proto/post-service --go-grpc_opt=paths=source_relative \
      --grpc-gateway_out=services/post-service/proto/post-service --grpc-gateway_opt=paths=source_relative \
      --openapiv2_out=services/post-service/docs --openapiv2_opt=allow_merge=true \
      services/post-service/api/post-service/*.proto

vendor-proto:
		@if [ ! -d vendor-protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor-protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor-protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi