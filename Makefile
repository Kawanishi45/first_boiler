run:
	go run ./cmd/server/main.go

sqlboiler:
	sqlboiler mysql -c sqlboiler.toml -o ./cmd/server/models --no-tests