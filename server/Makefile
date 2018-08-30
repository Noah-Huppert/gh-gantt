.PHONY: run test db

DB_NAME=dev-gh-gantt
DB_USERNAME=dev-gh-gantt
DB_DATA_DIR=${PWD}/run-data

# run server
run:
	go run main.go

# test source code
test:
	go test ./...

# db starts a local database
db:
	docker run \
		-it \
		--rm \
		--net host \
		-v ${DB_DATA_DIR}:/var/lib/postgresql/data \
		-e POSTGRES_DB=${DB_NAME} \
		-e POSTGRES_USER=${DB_USERNAME} \
		postgres
