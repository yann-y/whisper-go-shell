build:
	rm -rf dist/whisper-go
	GOOS=linux GOARCH=amd64 go build  -o dist/whisper-go  main.go

docker:
	sudo docker build --network=host -t  yanfive/whisper-go .