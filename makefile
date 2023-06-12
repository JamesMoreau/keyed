build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: clean build 
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build
	./keyed

clean:
	rm web/app.wasm
	rm keyed