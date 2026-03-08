build:
	go build -o hotreload ./cmd/hotreload

run:
	go run cmd/hotreload/main.go \
	--root . \
	--build "go build -o ./bin/server ./testserver" \
	--exec "./bin/server"

demo:
	./hotreload \
	--root . \
	--build "go build -o ./bin/server ./testserver" \
	--exec "./bin/server"
