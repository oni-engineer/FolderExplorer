.DEFAULT_GOAL := build

fmt:
	go fmt FolderExplorer.go
.PHONY:fmt

lint: fmt
	golint FolderExplorer.go
.PHONY:lint

vet: fmt
	go vet FolderExplorer.go
.PHONY:vet

build: vet
	go build FolderExplorer.go
.PHONY:build