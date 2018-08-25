.PHONY: dev db

DB_NAME=dev-gh-gantt

# dev runs the application for development
dev:
	buffalo dev

# db starts a local postgres database
db:
	mkdir -p run-data
	docker run \
		-it \
		--rm \
		--net host \
		-v ${PWD}/run-data:/var/lib/postgresql \
		-e POSTGRES_DB=${DB_NAME} \
		postgres
