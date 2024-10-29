tmpl:
	@templ generate -watch -proxy=http://localhost:42069

tailwind:
	@tailwindcss  -i ./css/tailwind.css -o css/main.css --watch

run:
	@npm run build
	@templ generate
	@go run cmd/main.go
