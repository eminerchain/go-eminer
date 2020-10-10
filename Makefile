# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: em android ios em-cross swarm evm all test clean
.PHONY: em-linux em-linux-386 em-linux-amd64 em-linux-mips64 em-linux-mips64le
.PHONY: em-linux-arm em-linux-arm-5 em-linux-arm-6 em-linux-arm-7 em-linux-arm64
.PHONY: em-darwin em-darwin-386 em-darwin-amd64
.PHONY: em-windows em-windows-386 em-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

em:
	build/env.sh go run build/ci.go install ./cmd/em
	@echo "Done building."
	@echo "Run \"$(GOBIN)/em\" to launch em."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

em-cross: em-linux em-darwin em-windows em-android em-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/em-*

em-linux: em-linux-386 em-linux-amd64 em-linux-arm em-linux-mips64 em-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-*

em-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/em
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep 386

em-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/em
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep amd64

em-linux-arm: em-linux-arm-5 em-linux-arm-6 em-linux-arm-7 em-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep arm

em-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/em
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep arm-5

em-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/em
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep arm-6

em-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/em
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep arm-7

em-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/em
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep arm64

em-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/em
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep mips

em-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/em
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep mipsle

em-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/em
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep mips64

em-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/em
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/em-linux-* | grep mips64le

em-darwin: em-darwin-386 em-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/em-darwin-*

em-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/em
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/em-darwin-* | grep 386

em-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/em
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/em-darwin-* | grep amd64

em-windows: em-windows-386 em-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/em-windows-*

em-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/em
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/em-windows-* | grep 386

em-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/em
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/em-windows-* | grep amd64
