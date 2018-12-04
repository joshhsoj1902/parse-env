build-docker:
	docker build -f Dockerfile -t joshhsoj1902/parse-env .

build:
	go build -o main cmd/main/main.go

run:
	./main --env "GO"

start:
	docker run joshhsoj1902/parse-env

.PHONY: build
