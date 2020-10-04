build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t demo-cli -f env/Dockerfile .
run:
	docker run demo-cli