API_PATH 		= api/razielsd/demo/v1
PROTO_API_DIR 	= api/razielsd/demo/v1
PROTO_OUT_DIR 	= server/pkg/razielsd/demo/v1
PROTO_API_OUT_DIR = ${PROTO_OUT_DIR}

.PHONY: gen-proto
proto:
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH} \
		-I third_party/googleapis \
		--include_imports \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
		--descriptor_set_out=$(PROTO_API_OUT_DIR)/api.pb \
        --go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		./${PROTO_API_DIR}/*.proto

.PHONY: lint
lint:
	cd template && golangci-lint run ./...
	cd builder && golangci-lint run ./...