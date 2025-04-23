BINARY_NAME=curlfs
SRC=main.go
OUTPUT_DIR=build

PLATFORMS = \
    linux/amd64 \
    linux/arm64 \
    windows/amd64 \
    darwin/amd64 \
    darwin/arm64

all: clean build

build:
	@mkdir -p $(OUTPUT_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*} GOARCH=$${platform#*/} \
		output_name=$(OUTPUT_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}; \
		if [ "$${platform%/*}" = "windows" ]; then output_name=$${output_name}.exe; fi; \
		echo "Building $$platform -> $$output_name"; \
		GOOS=$${platform%/*} GOARCH=$${platform#*/} go build -o $$output_name $(SRC); \
	done

clean:
	@rm -rf $(OUTPUT_DIR)

.PHONY: all build clean