
PROJECT=Luka
GOPATH ?= $(shell go env GOPATH)
CURDIR := $(shell pwd)

FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO        := GO111MODULE=on go
GOBUILD   := $(GO) build
GOTEST    := $(GO) test -p 4

FILES     := $$(find . -name "*.go")

default: keeperD

keeperD:
	@echo "generate keeper"
	$(GOBUILD) -o bin/KeeperDeployment/KeeperDeployment main/KeeperDeployment.go

fmt:
	@echo "gofmt (simplify)"
	@gofmt -s -l -w $(FILES) 2>&1 | $(FAIL_ON_STDOUT)
