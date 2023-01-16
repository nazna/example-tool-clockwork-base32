.PHONY: dev
dev:
	@$(MAKE) build-wasm
	@go run ./server.go

.PHONY: build
build:
	@$(MAKE) build-exec
	@$(MAKE) build-wasm

.PHONY: build-exec
build-exec:
	@go build -o ./dist/cwbase32 -ldflags="-s -w -X main.Version=0.0.999" -trimpath

.PHONY: build-wasm
build-wasm:
	@mkdir -p ./pages
	@cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js ./pages/wasm_exec_tinygo.js
	@cp ./index.html ./pages/index.html
	@tinygo build -o ./pages/cwbase32.wasm -target wasm -no-debug ./main.go

.PHONY: clean
clean:
	@rm -rf ./dist ./pages
