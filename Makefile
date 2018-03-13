.PHONY: run \
	redis

MAIN_FILE=main.go

REDIS_NAME=redis
REDIS_DIR=run-data

# run starts the Go HTTP server
run:
	go run "${MAIN_FILE}"

# redis starts a Redis server
redis:
	docker run \
		--name "${REDIS_NAME}" \
		-it \
		--rm \
		--net host \
		-p 6379:6379 \
		-v "$(shell pwd)/${REDIS_DIR}" \
		redis:alpine
