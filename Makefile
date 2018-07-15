.PHONY: run

MAIN_FILE=main.go

# run starts the Go HTTP server
run:
	go run "${MAIN_FILE}"
