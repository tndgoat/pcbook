PROTO_DIR=proto
PB_DIR=pb

PROTO_FILES=$(shell find $(PROTO_DIR) -name "*.proto")

# Tools
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

# =====================
# Generate protobuf code
# =====================
proto:
	@if [ -z "$(PROTOC_GEN_GO)" ]; then echo "Missing protoc-gen-go"; exit 1; fi
	@if [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then echo "Missing protoc-gen-go-grpc"; exit 1; fi

	mkdir -p $(PB_DIR)

	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(PB_DIR) \
		--go-grpc_out=paths=source_relative:$(PB_DIR) \
		$(PROTO_FILES)

# =====================
# Clean generated files
# =====================
clean:
	rm -rf $(PB_DIR)

# =====================
# Run
# =====================
run:
	go run main.go

.PHONY: proto clean run