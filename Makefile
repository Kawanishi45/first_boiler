run:
	go run ./cmd/server/main.go

sqlboiler:
	sqlboiler mysql -c sqlboiler.toml -o ./models --no-tests