all: 386 arm arm64

386:
	GOOS=linux GOARCH=386 go build -o bin/input2hid-linux-386

arm:
	GOOS=linux GOARCH=arm go build -o bin/input2hid-linux-arm

arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/input2hid-linux-arm64

