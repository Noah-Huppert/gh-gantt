# Server
GH Gantt server.

# Table Of Contents
- [Overview](#overview)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Run](#run)

# Overview
Provides the GH Gantt API and serves frontend assets.

# Configuration
Application configuration is provided via environment variables.  

See the [`config/config.go`](config/config.go) file for more information.

# Dependencies
Install Go dependencies with [Dep](https://golang.github.io/dep/):

```
dep ensure
```

# Run
## Database
Start a local Postgres database:

```
make db
```

See the [Run Book](RUN-BOOK.md) for details on how to run database migrations.

## Server
Start the server:

```
make run
```
