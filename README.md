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
Configuration parameters are passed via environment variables. See the 
`server/config/config.go` file for more information.

# Development
## Dependencies
### Server
[Dep](https://golang.github.io/dep/) is used to manage server Go dependencies.  

Install / update dependencies:

```
cd server
dep ensure
```

### Frontend
NPM is used to manage frontend NodeJS dependencies.  

Install / update dependencies:

```
cd frontend
npm install
```

## Start
### Server
Start the server:

```
cd server
make run
```

### Frontend Build
Build the frontend web app:

```
cd frontend
npm run build
```

Rebuild the frontend web app on changes by running:

```
cd frontend
make watch
```

This requires the 
[onchanges script](https://github.com/Noah-Huppert/scripts/blob/master/onchanges) 
to be in your path.
