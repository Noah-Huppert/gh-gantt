# Development 
GH Gantt development guide.

# Table Of Contents
- [Run Tests](#run-tests)
- [Writing a DB Migration](#writing-a-db-migration)

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
