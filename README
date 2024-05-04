# Go WASM module for Obsidian
My purpose is trying to use go wasm in Obsidian plugin.

## Enable WASM in go and Obsidian

### 1. compile WASM file
`make wasm` will generate `md-monitor.wasm` file which is executed in Obsidian plugin.

### 2. execute WASM file
There are two ways to execute wasm file:
1. browser(Chrome);
2. nodejs(Obsidian console or other electron app);

To execute wasm file in browser, please reference the file `test/index.html`. For nodejs, please reference the file `test/test_in_console.js`.

### 3. call exported function of wasm module
- export function:
```go
js.Global().Set("jsPI", jsPI())
```

- call exported function:
```javascript
WebAssembly.instantiate(content, go.importObject).then((ret) => { go.run(ret.instance); });
// call exported function which define in go module.
jsPI(3);
```

### 4. call js API in wasm
```go
js.Global().Call("alert", "this is an alerting!")
v := js.Global().Get("app")
fmt.Println(v.Get("title").String())
fmt.Println(v.Call("getAppTitle", "").String())
```

## TODO
- WASI