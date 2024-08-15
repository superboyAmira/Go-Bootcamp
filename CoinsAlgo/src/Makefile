PACKAGE_NAME = internal/minCoins

test:
	go test -v ./$(PACKAGE_NAME) -cover

bench:
	go test -bench=. ./$(PACKAGE_NAME) 

top10:
	go run cmd/day07/main.go
	go tool pprof -text cpu.prof | grep -A 10 "Showing nodes accounting" > top10.txt

doc:
	go install go101.org/golds@latest
	export PATH=$PATH:$HOME/go/bin
	golds -gen -dir=generated -nouses ./internal/minCoins ./cmd/day07
	zip -r docs.zip generated/

