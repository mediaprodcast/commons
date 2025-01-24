PROTO_DIR = ../protos
OUT_DIR = ./genproto
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
PROTO_NAMES = $(notdir $(basename $(PROTO_FILES))) # Extract proto file names without extensions

# Targets
.PHONY: all clean

all: $(PROTO_NAMES)

# Rule to generate for each proto file
$(PROTO_NAMES): %: $(PROTO_DIR)/%.proto
	mkdir -p $(OUT_DIR)/$@
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR)/$@ --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR)/$@ --go-grpc_opt=paths=source_relative \
		$<

clean:
	rm -rf $(OUT_DIR)
