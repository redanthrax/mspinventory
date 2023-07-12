build:
	@go build -o bin/app ./cmd/app/

run: build
	@./bin/app

clean:
	@rm -rf bin
