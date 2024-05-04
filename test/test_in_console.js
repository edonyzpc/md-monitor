'use strict'
// Try to execute the following code in browser or electron(Obsidian) console

const exec_wasm = require('/Users/edonymurphy/code/md-monitor/test/wasm_exec.js');
const go = new Go();
const fs = require('fs');
let content = undefined;

fs.readFile("/Users/edonymurphy/code/md-monitor/bin/md-monitor.wasm", (err, data) => { content = data });

WebAssembly.instantiate(content, go.importObject).then((ret) => { go.run(ret.instance); });

// call exported function which define in go module.
jsPI(3);