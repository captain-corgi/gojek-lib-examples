all: build run clean

build:
	@go build -o heimdall-example cmd/heimdall/main.go
run:
	@./heimdall-example
clean:
	@rm heimdall-example
