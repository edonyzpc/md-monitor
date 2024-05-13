.PHONY: wasm test

wasm:
	GOOS=js GOARCH=wasm go build -o  ./bin/md-monitor.wasm

test:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./test/
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec_node.js" ./test/

server:
	mkdir -p ./bin
	go build -o ./bin/server ./http-server/server.go