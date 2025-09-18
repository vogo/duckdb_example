docker-version:
	docker run --rm golang-cpp:1.25-alpine sh -c 'gcc --version'

build:
	docker run --rm -v "$$(pwd)":/app \
	-v ~/go:/go \
	go-cpp-builder:1.25 sh -c 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/main main.go'
