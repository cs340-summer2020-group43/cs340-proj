all: clean build

build: clean
	go build -o ./build/app ./src/app/app.go

clean:
	rm -f ./build/*
