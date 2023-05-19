#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
GO_MODULE := $(shell cat go.mod | grep module | cut -d ' ' -f 2)
GO_VERSION := $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2)
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf

GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
ENABLE_VERSION_CHECK ?= false

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(DOXCHAIN_BUILD_OPTIONS)))
  build_tags += gcc
else ifeq (rocksdb,$(findstring rocksdb,$(DOXCHAIN_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=doxchain \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=doxchaind \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq (cleveldb,$(findstring cleveldb,$(DOXCHAIN_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(DOXCHAIN_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
ifeq (,$(findstring nostrip,$(DOXCHAIN_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(DOXCHAIN_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

###############################################################################
###                                  Build                                  ###
###############################################################################
check_version:
ifeq ($(ENABLE_VERSION_CHECK),true)
	ifneq ($(GO_MINOR_VERSION),18)
		@echo "ERROR: Go version 1.18 is required for this version of Doxchain."
		exit 1
	endif
endif

build: check_version go.sum
	mkdir -p $(BUILDDIR)/
	go build -mod=readonly  $(BUILD_FLAGS) -o $(BUILDDIR)/ $(GO_MODULE)/cmd/doxchaind

build-release: build-release-amd64 build-release-arm64

build-release-amd64: go.sum $(BUILDDIR)/
	$(DOCKER) buildx create --name core-builder || true
	$(DOCKER) buildx use core-builder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
    --build-arg BUILDPLATFORM=linux/amd64 \
    --build-arg GOOS=linux \
    --build-arg GOARCH=amd64 \
		-t core:local-amd64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f core-builder || true
	$(DOCKER) create -ti --name core-builder core:local-amd64
	mkdir -p $(BUILDDIR)/release/
	$(DOCKER) cp core-builder:/usr/local/bin/doxchaind $(BUILDDIR)/release/doxchaind
	tar -czvf $(BUILDDIR)/release/doxchain_$(VERSION)_Linux_x86_64.tar.gz -C $(BUILDDIR)/release/ doxchaind
	rm $(BUILDDIR)/release/doxchaind
	$(DOCKER) rm -f core-builder

build-release-arm64: go.sum $(BUILDDIR)/
	$(DOCKER) buildx create --name core-builder || true
	$(DOCKER) buildx use core-builder 
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
    --build-arg BUILDPLATFORM=linux/arm64 \
    --build-arg GOOS=linux \
    --build-arg GOARCH=arm64 \
		-t core:local-arm64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f core-builder || true
	$(DOCKER) create -ti --name core-builder core:local-arm64
	$(DOCKER) cp core-builder:/usr/local/bin/doxchaind $(BUILDDIR)/release/doxchaind 
	tar -czvf $(BUILDDIR)/release/doxchain_$(VERSION)_Linux_arm64.tar.gz -C $(BUILDDIR)/release/ doxchaind 
	rm $(BUILDDIR)/release/doxchaind
	$(DOCKER) rm -f core-builder

install: check_version go.sum
	go install -mod=readonly $(BUILD_FLAGS) $(GO_MODULE)/cmd/doxchaind

###############################################################################
###                                Protobuf                                 ###
###############################################################################

CONTAINER_PROTO_VER=v0.7
CONTAINER_PROTO_IMAGE=tendermintdev/sdk-proto-gen:$(CONTAINER_PROTO_VER)
CONTAINER_PROTO_FMT=cosmos-sdk-proto-fmt-$(CONTAINER_PROTO_VER)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(CONTAINER_PROTO_IMAGE) sh ./scripts/protocgen.sh

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${CONTAINER_PROTO_FMT}$$"; then docker start -a $(CONTAINER_PROTO_FMT); else docker run --name $(CONTAINER_PROTO_FMT) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./proto -name "*.proto" -exec clang-format -i {} \; ; fi

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against '$(HTTPS_GIT)#branch=main'

.PHONY: proto-all proto-gen proto-format proto-lint proto-check-breaking 