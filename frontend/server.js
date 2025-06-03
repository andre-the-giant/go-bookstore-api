const http = require('http');
const fs = require('fs');
const path = require('path');

const PORT = 3000;
const API_HOST = 'api';
const API_PORT = 8989;

const MIME_TYPES = {
  '.html': 'text/html',
  '.js': 'application/javascript',
  '.css': 'text/css',
  '.json': 'application/json',
  '.png': 'image/png',
  '.jpg': 'image/jpeg',
  '.svg': 'image/svg+xml'
};

function serveStatic(filePath, res) {
  fs.readFile(filePath, (err, content) => {
    if (err) {
      res.writeHead(404, { 'Content-Type': 'text/plain' });
      res.end('Not found');
      return;
    }
    const ext = path.extname(filePath);
    const type = MIME_TYPES[ext] || 'text/plain';
    res.writeHead(200, { 'Content-Type': type });
    res.end(content);
  });
}

const server = http.createServer((req, res) => {
  if (req.url.startsWith('/api')) {
    const options = {
      hostname: API_HOST,
      port: API_PORT,
      path: req.url.replace('/api', ''),
      method: req.method,
      headers: req.headers
    };

    const proxy = http.request(options, (pRes) => {
      res.writeHead(pRes.statusCode, pRes.headers);
      pRes.pipe(res, { end: true });
    });

    req.pipe(proxy, { end: true });
    proxy.on('error', () => {
      res.writeHead(500);
      res.end('Proxy error');
    });
    return;
  }

  const filePath = path.join(__dirname, req.url === '/' ? 'index.html' : req.url);
  serveStatic(filePath, res);
});

server.listen(PORT, () => {
  console.log(`Frontend running on http://localhost:${PORT}`);
});
