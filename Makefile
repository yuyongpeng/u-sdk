GO=go
BUILD_DIR=build
ProjectPKGName=u-sdk

all: build-cbac-cmd

build:
	$(GO) build -ldflags "-s -w" -o $(BUILD_DIR)/cbac ./

clean:
	rm -f $(BUILD_DIR)/*


.PHONY: all build build-cbac-cmd
