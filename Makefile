.PHONY: app

app:
	go build -o build/pb-darwin
	env GOOS=linux GOARCH=amd64 go build -o build/pb-amd64
	env GOOS=linux GOARCH=386 go build -o build/pb-386
	env GOOS=windows GOARCH=amd64 go build -o build/pb-win

default: app