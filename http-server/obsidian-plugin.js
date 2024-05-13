'use strict'
const exec = require('child_process');
const http = require('http');


const options = {
    socketPath: '/Users/edony/code/md-monitor/bin/server.sock',
    path: '/',
};

const callback = res => {
    console.log(`STATUS: ${res.statusCode}`);
    res.setEncoding('utf8');
    res.on('data', data => console.log(data));
    res.on('error', data => console.error(data));
};

/*
exec.exec('/Users/edony/code/md-monitor/http-server/server', (err, stdout, stderr) => {
    if (err) {
        console.error(`exec error: ${err}`);
        return;
    }

    console.log(`Number of files ${stdout}`);
});
*/

const server = exec.spawn("/Users/edony/code/md-monitor/bin/server", [], {
    detached: false
});

server.on('error', (err) => {
    console.error(`exec error: ${err}`);
});

server.on('exit', (code) => {
    console.log(`child process exited with code ${code}`);
    process.exit(code);
})

server.stdout.on('data', (data) => {
    console.log(`stdout: ${data}`);
});

// sending request
const clientRequest = http.request(options, callback);
clientRequest.end();

// kill child process
// server.kill(9)
process.on("exit", () => child.kill())