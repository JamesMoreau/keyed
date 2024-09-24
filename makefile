wasmBuild:
	GOARCH=wasm GOOS=js go build -o main.wasm password.go

serverRun:
	python3 -m http.server

clean:
	rm -f githubpages/web/app.wasm
