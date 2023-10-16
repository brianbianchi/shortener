all: clean init serve

serve:
	go build -o bin/web main.go
	./bin/web

init:
	go build -o bin/init db/init.go
	./bin/init

clean:
	rm -rf bin

test:
	go test ./...