# React-Go Graphics Engine (scaffold)

This workspace contains a small scaffold demonstrating a WebGL renderer in a React + TypeScript frontend and a Go backend that streams angle updates over WebSocket and exposes tiny math/cipher utilities.

Quick start

- Start Go server:

```bash
cd go-server
go mod tidy
go run .
```

- Start frontend (in separate terminal):

```bash
cd javascript/frontend
npm install
npm run dev
```

Open http://localhost:3000 and the page will connect to the Go server at ws://localhost:8080/ws.

Files of interest

- [javascript/frontend](javascript/frontend) - React + Vite app with WebGL renderer.
- [go-server](go-server) - Go WebSocket server and helper packages `mathutils` and `cipher`.

Notes

- The Go server uses `github.com/gorilla/websocket`; `go mod tidy` will fetch it.
- The WebGL renderer is a tiny demo (rotating triangle) to show end-to-end data flow. Expand with more complex geometry, shader pipelines, and a math backend as desired.
