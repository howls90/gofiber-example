.PHONY: clean test build run docs swag pprof

APP_NAME = fiber

clean:
	rm -rf ./build

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o target/$(APP_NAME) main.go

run: build
	./target/$(APP_NAME)

watch:
	air

swag:
	swag init ./..

docs:
	godoc -http=:6060

pprof: build
	./target/$(APP_NAME) -performance

graphviz:
	go tool pprof -http=":8081" http://localhost:3000/debug/pprof/profile