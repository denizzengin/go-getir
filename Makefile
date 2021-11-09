build:
	go build -o bin/go-getir .
	go vet
	go fmt
	golint

run:
	go run  .

compile:
	echo "Compiling for every OS and Platform"
	
