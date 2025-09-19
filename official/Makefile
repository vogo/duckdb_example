docker-version:
	docker run --rm golang-cpp:1.25-alpine sh -c 'gcc --version'

build:
	docker run --rm -v "$$(pwd)":/app \
	-v ~/go:/go \
	golang-cpp:1.25-bookworm sh -c 'GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/main main.go'


build-static:
	docker run --rm -v "$$(pwd)":/app \
	-v ~/go:/go \
	go-cpp-builder:1.25 sh -c 'GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CPPFLAGS="-DDUCKDB_STATIC_BUILD" CGO_LDFLAGS="-lduckdb -lstdc++ -lm -ldl" go build -tags=duckdb_use_static_lib -ldflags "-s -w" -o dist/main main.go'