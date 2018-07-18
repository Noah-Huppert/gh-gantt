# GitHub Gantt
GitHub issues gantt chart generator.  

# Table Of Contents
- [Our Sponsors](#our-sponsors)
- [Overview](#overview)
- [Configuration](#setup)
- [Development](#development)

# Our Sponsors
[![Enzyme Logo](img/enzyme_logo_blue.svg)](https://www.enzyme.com)  
  
[Enzyme](https://www.enzyme.com) sponsors the development of this project.  

# Overview
GH Gantt displays GitHub issues in gantt chart form. ZenHub is used to augment 
GitHub issues with additional information.  

See the [wiki](https://github.com/Noah-Huppert/gh-gantt/wiki) for design information.

# Configuration
Make a copy of `config.ex.toml` named `config.toml` and fill it in with your 
own values.  

- HTTP: Web server configuration
	- Port: Port to handle HTTP traffic from
- RPC: Remote procedure call configuration
	- Port: Port to list for RPC calls on
- GitHub: Configuration related to GitHub issues API
	- AccessToken: GitHub API access token used to retrieve repo issues
	- RepoOwner: Login of GitHub user who owns repository
	- RepoName: Name of GitHub repository to retrieve issues from
- ZenHub:
	- APIToken: ZenHub API access token  
	            Must retrieve special ZenHub authentication token by:  
		- [Navigating to app.zenhub.com](https://app.zenhub.com)  
		- Run in the console  
		  ```js
		  window.localStorage.getItem("api_token")
		  ```

# Development
## Dependencies
### Server
[Dep](https://golang.github.io/dep/) is used to manage server Go dependencies.  

Install / update dependencies:

```
dep ensure
```

### Frontend
Pub is used to manage frontend Dart dependencies.  

Install / update dependencies:

```
cd frontend
pub get
```

TODO: Design frontend dev tools for watching frontend src and rebuilding
TODO: Get basic GRPC setup for frontend

## Protocol Buffers
[Protocol Buffers](https://developers.google.com/protocol-buffers/) is used to 
generate RPC code.  

Generate code from protocol buffers:

```
make proto
```

## Start Server
Start the server:

```
make run
```
