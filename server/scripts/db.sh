#!/usr/bin/env bash
#
# db.sh - Starts a local PostgreSQL server
#
# USAGE
#	db.sh [ENV] [--no-tty]
#
# ARGUMENTS
#	1. ENV    (Optional) Environment to start local server for. Value cannot be "prod". Defaults to "dev".
#
# OPTIONS
#	--no-tty    Starts the database without a TTY
#
# BEHAVIOR
#	Starts a local PostgreSQL server with a user and database named ENV-gh-gantt
#	Saves data in the run-data/ENV directory.

# Load arguments
opt_no_tty="false"

while [ ! -z "$1" ]; do
	if [[ "$1" == "--no_tty" ]]; then
		opt_no_tty="true"
	else
		db_env="$1"
		if [ -z "$db_env" ]; then
			db_env="dev"
		fi
	fi

	shift
done

if [[ "$db_env" == "prod" ]]; then
	echo "Error: ENV argument value cannot be \"prod\"" >&2
	exit 1
fi

# Configuration
db_data_dir="$PWD/$(dirname $0)/../run-data/$db_env"
db_name="$db_env-gh-gantt"
db_username="$db_env-gh-gantt"

echo "#######################"
echo "# Starting PostgreSQL #"
echo "#######################"
echo "Env           : $db_env"
echo "Data directory: $db_data_dir"
echo "DB Name       : $db_name"
echo "DB Username   : $db_username"

# Run
#mkdir -p "$db_data_dir"

run_args="i"

if [[ "$opt_no_tty" == "true" ]]; then
	run_args="${run_args}t"
fi

docker run \
	-$run_args \
	--rm \
	--net host \
	-v $db_data_dir:/var/lib/postgresql/data \
	-e POSTGRES_DB=$db_name \
	-e POSTGRES_USER=$db_username \
	postgres
