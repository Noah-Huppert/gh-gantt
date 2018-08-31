# Development 
GH Gantt development guide.

# Table Of Contents
- [Run API Server](#run-api-server)
- [Run Local Database](#run-local-database)
- [Configure](#configure)
- [Install Dependencies](#install-dependencies)
- [Run Tests](#run-tests)
- [Writing a DB Migration](#writing-a-db-migration)

# Run API Server
First [start a local database](#run-local-database).  

Then [set configuration values](#configure).  

Next [install dependencies](#install-dependencies).  

Finally start the server:

```
make run
```

# Run Local Database
The `db` Make target starts a local PostgreSQL server.  

The environment to start the database for can be passed as the first command line argument. If not provided the 
environment defaults to `dev`.

Usage: `make db [ENV=ENV]`  

Examples:

- `make db`
- `make db ENV=test`

*Note: `prod` cannot be passed as a valid `ENV` value*

# Configure
Application configuration is provided via environment variables.  

See the [`config/config.go`](config/config.go) file for more information.

# Install Dependencies
Install Go dependencies with [Dep](https://golang.github.io/dep/):

```
dep ensure
```

# Run Tests
Execute tests by running:

```
make test
```

# Writing a DB Migration
Database migrations are stored in the `server/migrations` directory.  

Each migration has a number prefixing it which indicates the order it will be applied.  

Each migration is composed up an "up" and "down" file. The "up" file has the `.up.sql` file extension. And contains an
SQL statement to execute the migration. The "down" file has the `.down.sql` file extension. And contains an SQL
statement to undo the migration.

To create a new migration make a `.up.sql` and `.down.sql` file. Prefix each file with the same number indicating the 
order the migration should be applied. After the number in the file name add a short description of what the migration
does.

Ex:

`1_create_users_table.up.sql`  
`1_create_users_table.down.sql`
