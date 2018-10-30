# Development 
GH Gantt development guide.

# Table Of Contents
- [Run API Server](#run-api-server)
- [Configure](#configure)
- [Install Dependencies](#install-dependencies)
- [Run Tests](#run-tests)

# Run API Server
First [set configuration values](#configure).  

then [install dependencies](#install-dependencies).  

Finally start the server:

```
make run
```

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
