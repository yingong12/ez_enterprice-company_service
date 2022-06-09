os?=linux


run: 
	go run main.go  bootstrap.go 

build:export GOOS=$(os)
build:export GOARCH=amd64
build:
	@echo "building binary for $(GOOS)..."
	go build -o ./company_service 
	@echo "done!"
	