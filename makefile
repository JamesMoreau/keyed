build_wasm:
	GOARCH=wasm GOOS=js go build -o main.wasm password.go

server:
	python3 -m http.server

clean:
	rm -f githubpages/web/app.wasm
	rm -f keyed
