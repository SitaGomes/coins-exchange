run:
	echo "Running server..."
	go mod tidy
	go run cmd/server/main.go