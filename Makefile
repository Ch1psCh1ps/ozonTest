IMAGE_NAME := graphql-server
PORT := 8080

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p $(PORT):$(PORT) $(IMAGE_NAME)

start: build run

stop:
	docker stop $(shell docker ps -q --filter ancestor=$(IMAGE_NAME))

# Удаление
clean: stop
	docker rm $(shell docker ps -a -q --filter ancestor=$(IMAGE_NAME))
	docker rmi $(IMAGE_NAME)

# Линтер
lint:
	golangci-lint run --fix --timeout=5m -c .golangci.yml
