# Development 
GH Gantt development guide.

# Table Of Contents
- [Run API Server](#run-api-server)
- [Configure](#configure)
- [Generate GitHub State Signing Key](#generate-github-state-signing-key)
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

# Generate GitHub State Signing Key
The server uses a ed25519 keypair to sign a `state` field in GitHub authentication requests.  

Generate this key by running:

```
./scripts/gen-gh-state-signing-key.sh
```

This will save the key in the appropriate configuration environment variables in the `.env` file.

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
