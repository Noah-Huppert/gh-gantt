# GitHub Gantt
Active Work In Progress.  
GitHub issues gantt chart generator.

# Table Of Contents
- [Overview](#overview)
- [Setup](#setup)
- [Undocumented ZenHub API Endpoints](#undocumented-zenhub-api-endpoints)
	- [Dependences](#dependences)
# Overview
Creates a Gantt chart from a GitHub repositories issues. Uses ZenHub to 
retrieve the dependency information necessary to create a Gantt chart.  

The scheduling of tasks is determined by the Milestone they are in.  

Current status: Heavy development, screenshot:  

![Gantt chart generated with gh-gantt tool](/static/img/screenshot.png)

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

# Undocument ZenHub API Endpoints
ZenHub documents many of their endpoints. However some have not been documented, 
and provide key information.  

## Dependences
Provides information about an issue's dependencies.  

### Request
URL: `api.zenhub.io/v4/repositories/:repo_id/issues/:issue_number/dependencies`  

| Key            | Description                                     |
| -------------- | ----------------------------------------------- |
| `repo_id`      | ID of GitHub repository to retrieve issues from | 
| `issue_number` | Issue number to retrieve dependencies for       |

### Response
JSON payload with fields:

- `blocked_by` (`[]Dependency`): Array of dependency information about items 
				 currently blocking the specified issue
- `blocking` (`[]Dependency`): Array of dependency information about

Dependency information is returned in the form of a Dependency object. Which 
is a combination of the repository, issue, and pull request models.  

Example:  
```json
{
	"blocked_by": [
		{
			"issue_number":1,
			"repo_id":1234567,
			"cached_repo_name":"repo name",
			"cached_repo_owner":"repo owner",
			"updated_at":"2018-03-10T22:26:43Z",
			"closed_at":null,
			"created_at":"2018-03-10T22:26:05Z",
			"html_url":"https://github.com/repo owner/repo name/pull/1",
			"title":"issue title",
			"state":"open",
			"pull_request": {
				"patch_url":"https://github.com/repo owner/repo name/pull/1.patch",
				"diff_url":"https://github.com/repo owner/repo name/pull/1.diff",
				"html_url":"https://github.com/repo owner/repo name/pull/1",
				"url":"https://api.github.com/repos/repo owner/repo name/pulls/1"
			},
			"number":1,
			"milestone":null,
			"labels":[],
			"assignees":[],
			"assignee":null,
			"user":{"login":"a user"},
			"pipeline": {
				"name":"New Issues",
				"_id":"122345566789"
			},
			"blocked":true
		}
	],
	"blocking":[]
}
```
