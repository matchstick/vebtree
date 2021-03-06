# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: fix vet fmt test build tidy lint release

default: build

GOBIN := $(shell go env GOPATH)/bin
REPO_GOLINT := github.com/golangci/golangci-lint/cmd/golangci-lint
REPO_GOIMPORTS := golang.org/x/tools/cmd/goimports
REPO_GODOC := golang.org/x/tools/cmd/godoc

build:
	go build -v

cross:
	env GOOS=linux   GOARCH=amd64 go build -v -o vebtree.linux
	env GOOS=darwin  GOARCH=amd64 go build -v -o vebtree.darwin
	env GOOS=windows GOARCH=amd64 go build -v -o vebtree.exe

all: fix vet fmt tidy build lint test

fix:
	go fix ./...

fmt:
	go fmt ./...
	goimports -w .
	gofmt -w -s .

tidy:
	go mod tidy

test:
	go test -timeout 60m ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

docs:
	@echo "http://127.0.0.1:6060/pkg/#stdlib"
	@godoc -http=localhost:6060

covhtml:
	go test ./... -coverprofile=cov.out
	go tool cover -html=cov.out

covfunc:
	go test ./... -coverprofile=cov.out
	go tool cover -func=cov.out

clean:
	rm -f vebtree cov.out *.bak *.json \
		vebtree.darwin vebtree.exe vebtree.linux
