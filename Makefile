CURDIR := $(shell pwd)
GOBIN := $(CURDIR)/bin
MOCKDIR := $(CURDIR)/mmock

##
## List of commands:
##

## default:
all: mod deps test

all-deps: mod deps

deps:
	@echo "======================================================================"
	@echo 'MAKE: deps...'
	@mkdir -p $(GOBIN)
	@mkdir -p $(MOCKDIR)
	@GOBIN=$(GOBIN) GO111MODULE=on go get -v github.com/golang/mock/mockgen@latest

# Run test
test:
	go test --check.format=teamcity ./...

# Download all dependencies
mod:
	@echo "======================================================================"
	@echo "Run MOD"
# 	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod tidy -v
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod vendor -v
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod download
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify
	@echo "======================================================================"

clean-cache:
	@echo "clean-cache started..."
	go clean -cache
	go clean -testcache
	@echo "clean-cache complete!"
	@echo "clean-cache complete!"

mock:
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/rmq Queue > $(MOCKDIR)/queue_mock.go
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/rmq BatchConsumer > $(MOCKDIR)/batch_consumer_mock.go
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/rmq Consumer > $(MOCKDIR)/consumer_mock.go
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/rmq RedisClient > $(MOCKDIR)/redis_client_mock.go
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/rmq Connection > $(MOCKDIR)/connection_mock.go


