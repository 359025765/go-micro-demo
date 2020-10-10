build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t demo-cli -f .
run:
	docker run demo-cli