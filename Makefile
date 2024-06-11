run:
	@npm run build
	@templ generate
	@go run cmd/main.go
