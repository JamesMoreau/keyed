build:
	GOARCH=wasm GOOS=js go build -o githubpages/web/app.wasm
	go build

run: clean build 
	./keyed

clean:
	rm -f githubpages/web/app.wasm
	rm -f fkeyed
