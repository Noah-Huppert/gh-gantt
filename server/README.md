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

# Run
Install the Revel command line tool:

```
go get github.com/revel/cmd/revel
```

Run the server with the Revel command line tool:

```
make run
```
