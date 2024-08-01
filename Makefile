CURRENT_DIR := $(shell pwd)
WORK_DIR_BACKEND := $(CURRENT_DIR)/backend
WORK_DIR_FRONTEND := $(CURRENT_DIR)/front
BINARY_NAME := moments
COMMIT = $(shell git rev-parse --short HEAD)
LINUX_BINARY_NAME = $(BINARY_NAME)-linux-amd64
MACOS_BINARY_NAME = $(BINARY_NAME)-macos-amd64
WINDOWS_BINARY_NAME = $(BINARY_NAME)-windows-amd64.exe

export GO := go

.PHONY: clean build backend frontend


build: clean frontend backend

clean:
	$(GO) clean
	rm -rf $(WORK_DIR_BACKEND)/dist
	rm -rf $(WORK_DIR_BACKEND)/public
	rm -rf $(WORK_DIR_FRONTEND)/.output
	rm -rf $(WORK_DIR_FRONTEND)/dist

frontend:
	cd $(WORK_DIR_FRONTEND) && pnpm i && pnpm generate
	cp -r $(WORK_DIR_FRONTEND)/.output/public  $(WORK_DIR_BACKEND)/

backend:
	cd $(WORK_DIR_BACKEND) && CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GO) build -tags prod -ldflags="-s -w -X 'main.gitCommitID=$(COMMIT)'" -o $(WORK_DIR_BACKEND)/dist/$(MACOS_BINARY_NAME)
	cd $(WORK_DIR_BACKEND) && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GO) build -tags prod -ldflags="-s -w -X 'main.gitCommitID=$(COMMIT)'" -o $(WORK_DIR_BACKEND)/dist/$(LINUX_BINARY_NAME)
	cd $(WORK_DIR_BACKEND) && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GO) build -tags prod -ldflags="-s -w -X 'main.gitCommitID=$(COMMIT)'" -o $(WORK_DIR_BACKEND)/dist/$(WINDOWS_BINARY_NAME)
	cd $(WORK_DIR_BACKEND) && upx --best --lzma $(WORK_DIR_BACKEND)/dist/$(WINDOWS_BINARY_NAME)
	cd $(WORK_DIR_BACKEND) && upx --best --lzma $(WORK_DIR_BACKEND)/dist/$(LINUX_BINARY_NAME)
