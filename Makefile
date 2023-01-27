build:
	go build -o ./bin/pingvin-app ./cmd/pingvin-project/main.go
run: build
	./bin/pingvin-app --addr 127.0.0.1 --port 8081
clean:
	rm -rf bin/