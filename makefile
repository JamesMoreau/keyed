build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: clean build 
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build
	./keyed

clean:
	rm -f web/app.wasm
	rm -f fkeyed