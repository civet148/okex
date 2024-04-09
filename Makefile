#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`

build:
	rm -f okex
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o okex cmd/main.go
.PHONY: build
BINS+=okex

clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean
