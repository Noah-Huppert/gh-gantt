.PHONY: run proto clean-proto

# Configuration
MAIN_FILE=main.go

PROTO_SRC=repos/*.proto

# run starts the Go HTTP server
run:
	go run "${MAIN_FILE}"

# proto generates code from protocol buffers definitions
proto: clean-proto
	protoc ${PROTO_SRC} \
		--go_out=plugins=grpc:.

# clean-proto removes existing protocol buffers generated code
clean-proto:
	rm */*.pb.go || true
