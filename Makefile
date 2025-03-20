run:
	air

clean:
	@pkill -f "air" || true
	@pkill -f "main" || true

clean-css:
	@rm -rf public/styles


hash-css: watch-scss hash

watch-scss:
	@bunx sass --watch --load-path=node_modules src/styles/main.scss public/styles/main.css --style=compressed &

hash:
	@sleep 2 && go run ./cmd/hash/main.go

css:
	@bunx sass --load-path=node_modules src/styles/main.scss public/styles/main.css --style=compressed --silence-deprecation=import
	go run ./cmd/hash/main.go

watch-css:
	@bunx sass --watch --load-path=node_modules src/styles/main.scss public/styles/main.css --style=compressed 

build-css:
	@./bin/build-css


build-ts-scripts:
	@bun build ./dev/bun/build-css.ts --compile --outfile ./bin/build-css

docker-build:
	docker build -t goth-todo .

docker-run:
	docker run -p 8080:8080 goth-todo

docker-dev:
	docker compose up --build