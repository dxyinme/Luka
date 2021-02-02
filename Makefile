
PROJECT=Luka
GOPATH ?= $(shell go env GOPATH)
CURDIR := $(shell pwd)

FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO      	:= GO111MODULE=on GOARCH="amd64" go
GOARM		:= GO111MODULE=on GOARCH="arm" go
GOBUILD 	:= $(GO) build
GOARMBUILD 	:= $(GOARM) build
GOTEST  	:= $(GO) test -p 4

FILES   := $$(find . -name "*.go")

default: keeperD

keeperD:
	@echo "generate keeper"
	$(GOBUILD) -o bin/KeeperDeployment/KeeperDeployment main/KeeperDeployment.go
	@cp -rf conf bin/KeeperDeployment/
	@cp -rf script/keeper bin/KeeperDeployment/

assigneerD:
	@echo "generate assigneer"
	$(GOBUILD) -o bin/AssigneerDeployment/AssigneerDeployment main/AssigneerDeployment.go
	@cp -rf script/assigneer bin/AssigneerDeployment/
	@cp -rf conf bin/AssigneerDeployment/

assign-cli:
	@echo "generate assign-cli"
	$(GOBUILD) -o bin/assign-cli main/assign-cli.go

DBServer:
	@echo "generate DBServer"
	$(GOBUILD) -o bin/DBServer main/DBServer.go

AuthMain:
	@echo "generate AuthMain"
	$(GOBUILD) -o bin/AuthMain main/AuthMain.go

DBServerARM:
	@echo "generate DBServerARM"
	$(GOARMBUILD) -o arm_bin/DBServer_arm main/DBServer.go

keeperDARM:
	@echo "generate keeperARM"
	$(GOARMBUILD) -o arm_bin/KeeperDeployment/KeeperDeployment main/KeeperDeployment.go
	@cp -rf conf arm_bin/KeeperDeployment/
	@cp -rf script/keeper arm_bin/KeeperDeployment/

assigneerDARM:
	@echo "generate assigneerARM"
	$(GOARMBUILD) -o arm_bin/AssigneerDeployment/AssigneerDeployment main/AssigneerDeployment.go
	@cp -rf script/assigneer arm_bin/AssigneerDeployment/
	@cp -rf conf arm_bin/AssigneerDeployment/

assign-cliARM:
	@echo "generate assign-cliARM"
	$(GOARMBUILD) -o arm_bin/assign-cli main/assign-cli.go

AuthMainARM:
	@echo "generate AuthMainARM"
	$(GOARMBUILD) -o arm_bin/AuthMain main/AuthMain.go

fmt:
	@echo "gofmt (simplify)"
	@gofmt -s -l -w $(FILES) 2>&1 | $(FAIL_ON_STDOUT)

