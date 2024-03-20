server:
	nodemon --watch './**/*.go' --signal SIGTERM --exec APP_ENV=dev 'go' run cmd/main.go

swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/main.go