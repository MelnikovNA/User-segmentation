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
	./proto/user_segmentation/segment.proto 
	@echo "segmentation - done"