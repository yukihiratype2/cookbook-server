GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: output/cookbook-server

build-native: $(GOFILES)
	go build -o workdir/native-contacts .

output/cookbook-server: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o output/cookbook-server .
