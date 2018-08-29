# Run Book
GH Gantt Server run book.

# Table Of Contents
- [Database Migrations](#database-migrations)

# Database Migrations
> Instructions on how to run database migrations.

The `script/db-migrate.go` script runs database migrations.

Configuration is passed via the same database environment variables which are used to configure the API server. See the `config/db.go` file for details.  

Run the script by executing:

```
go run scripts/db-migrate.go
```


