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

See the [`config/config.go`](server/config/config.go) file for more information.

# Dependencies
Install Go dependencies with [Dep](https://golang.github.io/dep/):

```
dep ensure
```

# Run
Start the server:

```
make run
```
