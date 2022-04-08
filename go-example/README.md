This is a simple exmple of how to execute Go code on the web using Webassembly
The code here is based on this [Go Webassembly Tutorial](https://tutorialedge.net/golang/go-webassembly-tutorial/).

## Compile the .wasm
```bash
GOARCH=wasm GOOS=js go build -o lib.wasm hello.go
```

## Start the server
```bash
go run server.go
```

Visit the server on `localhost:8080`
