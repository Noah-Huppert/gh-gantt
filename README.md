# GitHub Gantt
GitHub issues gantt chart generator.  

# Our Sponsors
<a href="https://www.enzyme.com" target="_blank">
![Enzyme Logo](img/enzyme_logo_blue.svg)
</a>

# Table Of Contents
- [Overview](#overview)
- [Setup](#setup)

# Overview
GH Gantt displays GitHub issues in gantt chart form. ZenHub is used to augment 
GitHub issues with additional information.

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
