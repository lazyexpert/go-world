all:
	go build -o cmd/simulation/simulation internal/simulation/*go
	
dotest:
	./scripts/test

dorelease:
	./scripts/release