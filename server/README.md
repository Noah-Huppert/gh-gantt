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

See the `config/config.go` file for more information.

# Dependencies
Install Go dependencies with [Dep](https://golang.github.io/dep/):

```
dep ensure
```

Install the Buffalo command line tool:

```
go get github.com/gobuffalo/buffalo/buffalo
```


# Run
Start a local database:

```
make db
```

Start the server with the Buffalo command line tool:

```
make dev
```

The server will automatically restart when any changes to the source code are made.
