PB_REL := "https://github.com/protocolbuffers/protobuf/releases"
PROTOC_VERSION := 3.19.4

arch := linux-x86_64
ifeq (aarch64,$(shell uname -i))
	arch = linux-aarch_64
endif
ifeq (x86_32,$(shell uname -i))
	arch = linux-x86_32
endif

all: protoc generate tidy build

PROTOC = $(shell pwd)/bin/protoc
protoc:
	@[ -f $(PROTOC) ] || { \
	set -e ; \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26; \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1; \
	mkdir -p bin/; \
	TMP_DIR=$$(mktemp -d); \
	cd $$TMP_DIR ; \
	curl -LO $(PB_REL)/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(arch).zip; \
	unzip protoc-$(PROTOC_VERSION)-$(arch).zip -d .; \
	mv bin/protoc $(PROTOC); \
	rm -rf $$TMP_DIR ; \
	}

generate:
	$(PROTOC) --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/token/token.proto

tidy:
	go mod tidy -compat=1.17

build:
	go build -o bin/tokenmanager main.go

archive:
	git archive --format zip -o tokenmanager.zip master
