run:   
	go run ./cmd/app
build: 
	go build -o bin/app ./cmd/app
tidy:  
	go mod tidy