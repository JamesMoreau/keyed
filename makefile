build:
	GOARCH=wasm GOOS=js go build -o githubpages/web/app.wasm
	go build

run: clean build 
	./keyed

server:
	python3 -m http.server

clean:
	rm -f githubpages/web/app.wasm
	rm -f keyed
