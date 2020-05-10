
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
