# GitHub Gantt
GitHub issues gantt chart generator.  

# Our Sponsors
[![Enzyme Logo](img/enzyme_logo_blue.svg)](https://www.enzyme.com/)  

# Table Of Contents
- [Overview](#overview)
- [Setup](#setup)
- [Undocumented ZenHub API Endpoints](#undocumented-zenhub-api-endpoints)
	- [Dependences](#dependences)

# Overview
Creates a Gantt chart from GitHub repository issues. Uses ZenHub to 
retrieve the dependency information necessary to create a Gantt chart.  

The scheduling of tasks is determined by which Milestone they are in.  

Current status: Heavy development, screenshot:  

<img alt="Gantt chart generated with gh-gantt tool" width="400" src="/img/screenshot.png">

# Setup
## Configuration File
Make a copy of `config.ex.toml` named `config.toml` and fill it in with your 
own values.  

- HTTP: Web server configuration
	- Port: Port to handle HTTP traffic from
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
