build:
	@go build -o bin/app ./cmd/app/

run: build
	@./bin/app

clean:
	@rm -rf bin

tailwind:
	npx tailwindcss -i tailwind.css -o public/assets/app.css
