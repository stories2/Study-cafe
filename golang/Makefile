build:
	go build -o main.out
test:
	go test
run:
	./main.out
docker:
	docker build -t gostudy -f Dockerfile .
build-run: build run

docker-run:
	docker run --rm gostudy ./main.out