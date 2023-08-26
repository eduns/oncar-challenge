run-app:
	docker compose up --build

run-tests:
	cd backend && go test ./tests/**/*.go