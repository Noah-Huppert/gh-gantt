#!/usr/bin/env bash
#
# db.sh - Starts a local PostgreSQL server
#
# USAGE
#	db.sh [ENV]
#
# ARGUMENTS
#	1. ENV    (Optional) Environment to start local server for. Value cannot be "prod". Defaults to "dev".
#
# BEHAVIOR
#	Starts a local PostgreSQL server with a user and database named ENV-gh-gantt
#	Saves data in the run-data/ENV directory.

# Load arguments
db_env="$1"
if [ -z "$db_env" ]; then
	db_env="dev"
fi

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
mkdir -p "$db_data_dir"
docker run \
	-it \
	--rm \
	--net host \
	-v $db_data_dir:/var/lib/postgresql/data \
	-e POSTGRES_DB=$db_name \
	-e POSTGRES_USER=$db_username \
	postgres
