#!/usr/bin/env bash
#
# db.sh - Starts a local PostgreSQL server
#
# USAGE
#	db.sh [ENV]
#
# ARGUMENTS
#	1. ENV    (Optional) Environment to start local server for. Value cannot be "prod". Defaults to "dev".

# Configuration
db_data_dir="$PWD/$(dirname $0)/../run-data"

# Load arguments
db_env="$1"
if [ -z "$db_env" ]; then
	db_env="dev"
fi

if [[ "$db_env" == "prod" ]]; then
	echo "Error: ENV argument value cannot be \"prod\"" >&2
	exit 1
fi

# Run
db_name="$db_env-gh-gantt"
db_username="$db_env-gh-gantt"

echo "#######################"
echo "# Starting PostgreSQL #"
echo "#######################"
echo "Env        : $db_env"
echo "DB Name    : $db_name"
echo "DB Username: $db_username"

docker run \
	-it \
	--rm \
	--net host \
	-v $db_data_dir:/var/lib/postgresql/data \
	-e POSTGRES_DB=$db_name \
	-e POSTGRES_USER=$db_username \
	postgres
