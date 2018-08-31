#!/usr/bin/env bash
#
# test.sh - Runs tests on server source code
#
# USAGE
#	test.sh
#
# BEHAVIOR
#	Starts a local test PostgreSQL database server, runs server tests, and stops of the database server.

# Start database
echo "#####################"
echo "# Starting Database #"
echo "#####################"

db_log_f="/tmp/gh-gantt-test-db.out"

function show_db_logs() {
	# Show database logs on error
	echo "Database logs"
	echo "vvvvvvvvvvvvv"

	cat $db_log_f
}

cd "$(dirname $0)" && \
	./db.sh "test" --no-tty &> $db_log_f &

db_pid="$!"

# ... Wait for server to start
echo "Waiting for database to start"

started="false"
for i in $(seq 10); do
	# Check if process is running
	if ! kill -0 "$db_pid" &> /dev/null; then
		echo
		echo "Failed to start database" >&2

		show_db_logs

		exit 1
	fi

	# Check if database listening on port
	curl localhost:5432 &> /dev/null

	if [[ "$?" == "52" ]]; then
		started="true"
		break
	else
		printf "."
		sleep 1
	fi
done

if [[ "$started" == "true" ]]; then
	echo
	echo "Started"
else
	echo
	echo "Failed to start database" >&2

	show_db_logs

	exit 1
fi

# Run tests
echo "#################"
echo "# Running Tests #"
echo "#################"

src_dir="$PWD/$(dirname $0)/.."

cd "$src_dir" && \
	go test ./...

if [[ "$?" != "0" ]]; then
	echo "Failed to test server" >&2
else
	echo "Successfully tested server"
fi

# Stop database
echo "#####################"
echo "# Stopping Database #"
echo "#####################"


if kill -0 "$db_pid" &> /dev/null; then
	kill -KILL "$db_pid"
	echo "Stopped"
else
	echo "Not running"
fi

