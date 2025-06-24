BINARY=HTTP_IO_BOUND
PATH_TO_NAME=cmd/REST_IO_Bound/main.go
DOCKER_IMAGE_NAME=http_io_bound:v1

all: build_docker run_docker

test:
	@ echo "There should be tests here but they are in the process of being written"
run:
	@ go run $(PATH_TO_NAME)
build:
	@ CGO_ENABLED=0 go build -o $(BINARY) $(PATH_TO_NAME)
build_docker:
	@ docker build -t $(DOCKER_IMAGE_NAME) .
run_docker:
	@ docker run -it --rm -p 56789:8080 $(DOCKER_IMAGE_NAME)