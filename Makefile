templ:
	@npx tailwindcss -i view/css/input.css -o public/app.css #--watch
	@~/go/bin/templ generate --watch --proxy="http://localhost:3000" --cmd="go run ." --open-browser=false
	@go run .


install:
	@echo "Installing neccessary packages and dependencies..."
	@echo "Installing air for live reloading..."
	@go install github.com/air-verse/air@latest
	@echo "Installing Templ for templating..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@echo "Installing Node packages"
	@npm install
	@echo "Everything is done, to start developing type air"
