# Run Book
GH Gantt Server run book.

# Table Of Contents
- [Database Migrations](#database-migrations)
	- [Install Migration Tool](#install-migration-tool)
	- [Run Migrations](#run-migrations)

# Database Migrations
Instructions on how to run database migrations.

## Install Migration Tool
The [Migrate tool](https://github.com/golang-migrate/migrate) is used to run database migrations.  

1. Navigate to the [Migrate tool releases page](https://github.com/golang-migrate/migrate/releases) and download the 
	binary for your platform.  

2. Decompress the binary download:
   ```
   tar -xzf binary-download.tar.gz
   ```

3. Move the binary download to a location of your choice (On linux the `/opt/` directory is a good choice)  

4. Add the binary download to your `PATH` environment variable
	- On linux edit your shell profile file to include the install directory in your path
      ```
      export PATH="$PATH:/opt/migrate`
      ```
    - On Windows edit your system's environment variables in the settings

