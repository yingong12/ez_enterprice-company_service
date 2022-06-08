os?=linux
port?=8686


run: 
	swag init -g http/router.go
	go run main.go 

build:export GOOS=$(os)
build:export GOARCH=amd64
build:
	@echo "building binary for $(GOOS)..."
	go build -o ./company_service 
	@echo "done!"
	