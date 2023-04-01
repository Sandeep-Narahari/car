# Directories
PROTO_DIR := ./proto
PROTO_GEN_DIR := paths=source_relative:./types

# Proto files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Targets
.PHONY: proto-gen
proto-gen: generate

.PHONY: generate
generate: generate-grpc generate-gateway

.PHONY: generate-grpc
generate-grpc:
	protoc -I=$(PROTO_DIR) --go_out=$(PROTO_GEN_DIR) --go-grpc_out=$(PROTO_GEN_DIR) $(PROTO_FILES)

.PHONY: generate-gateway
generate-gateway:
	protoc -I=$(PROTO_DIR) --grpc-gateway_out=$(PROTO_GEN_DIR) --go-grpc_out=$(PROTO_GEN_DIR) $(PROTO_FILES)
