DEST=./codegen/go
DEST2= ./codegen/swagger
ERRORS_GENERATED := ./codegen/go/apierrors/apierrors_gen.go
ERRORS_SOURCE    := ./errors/errors.json

generate:
	@rm -rf codegen; 
	@mkdir -p ${DEST}/segmentation
	@mkdir -p  ${DEST}/common
	@mkdir -p ${DEST2}/segmentation 
	@mkdir -p  ${DEST2}/common
	@echo "start proto generation"; 

#common
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/common/common.proto
	@echo "common - done"

#user
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional=true \
    --openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/segmentation/segment.proto 
	@echo "segmentation - done"

#info
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/common/info.proto
	@echo "info - done"

	@echo "finish proto generation";

	@make gen


gen:
	protoc --proto_path=./proto \
	-I ./proto -I . -I ./proto/third_party\
    --descriptor_set_out=myservice.protoset \
    --include_imports ./proto/segmentation/segment.proto
	@echo "protoc - done"

errors-gen:
	go run cmd/err-gen/gen.go
	gofmt -w $(ERRORS_GENERATED)