all:
	make clean && mkdir bin && make compile
clean: 
	rm -rf bin
compile: 
	go build -o bin/world *.go
