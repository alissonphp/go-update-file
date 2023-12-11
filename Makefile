MAIN_PATH=cmd/go-update-file/main.go
GOOS=linux

default: build
build:
	@echo "Building binary application..."
	CGO_ENABLED=0 GOOS=$(GOOS) go build -ldflags "-s -w" -a -installsuffix cgo -o main $(MAIN_PATH)
