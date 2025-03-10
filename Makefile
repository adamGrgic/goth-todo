run:
	# @pkill -f "air" || true
	# @pkill -f "main" || true
	# sleep 1
	air

clean:
	@pkill -f "air" || true
	@pkill -f "main" || true

tailwind:
	bunx tailwindcss -c tailwind.config.js \
		-i ./client/src/main.css \
		-o ./client/dist/output.css \
		--watch

docker-build:
	docker build -t goth-todo .

docker-run:
	docker run -p 8080:8080 goth-todo

docker-dev:
	docker compose up --build