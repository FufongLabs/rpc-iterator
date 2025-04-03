PROJECT := $(shell basename "$(PWD)")
GO_BIN := $(CURDIR)/build

GOPROXY ?= "https://proxy.golang.org,direct"

.PHONY: all clean

all: iterate-rpc

iterate-rpc:
	GOPROXY=$(GOPROXY) \
	go build -ldflags "-s -w" \
	-o $(GO_BIN)/iterate-rpc \
	./cmd


clean:
	rm -fr ./build/*
