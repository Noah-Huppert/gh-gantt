.PHONY: run db test

# run server
run:
	go run main.go

# db runs a local PostgreSQL server
# Args:
#	- ENV: Environment to run database for, defaults to "dev"
db:
	./scripts/db.sh "${ENV}"

# test source code
test:
	./scripts/test.sh
