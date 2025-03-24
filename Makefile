run:
	air

check-env:
	@command -v docker >/dev/null || (echo "Docker is missing!"; exit 1)
	@command -v node >/dev/null || (echo "Node is missing!"; exit 1)

clean:
	@pkill -f "air" || true
	@pkill -f "main" || true

css:
	@./bin/css-compiler

css-watch:
	@./bin/css-compiler --watch

js:
	@./bin/javascript-compiler

js-watch:
	@./bin/javascript-compiler --watch

templ:
	@templ generate

templ-watch:
	@clear
	@templ generate --watch

docs-go: 
	@swag init -g ./cmd/app/main.go -o ./docs

bundler:
	@bun build ./src/scripts/bundlers/css.ts --compile --outfile ./bin/css-compiler
	@bun build ./src/scripts/bundlers/javascript.ts --compile --outfile ./bin/javascript-compiler

bundler-js-test:
	@bun ./src/scripts/bundlers/javascript.ts

bundler-css-test:
	@bun ./src/scripts/bundlers/css.ts

docker-build:
	docker build -t goth-todo .

docker-run:
	docker run -p 8080:8080 goth-todo

docker-dev:
	docker compose up --build