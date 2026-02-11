migrate:
	go run cmd/migration/main.go

run:
	go run cmd/web/main.go

windows-run:
	go build -o main.exe cmd/web/main.go
	main.exe
