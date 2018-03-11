# GitHub Gantt
GitHub issues gantt chart generator.

# Table Of Contents
- [Setup](#setup)
- [Undocumented ZenHub API Endpoints](#undocumented-zenhub-api-endpoints)
	- [Dependences](#dependencies)

# Setup
## Configuration File
Make a copy of `config.ex.toml` named `config.toml` and fill it in with your 
own values.  

- GitHub: Configuration related to GitHub issues API
	- AccessToken: GitHub API access token used to retrieve repo issues
	- RepoOwner: Login of GitHub user who owns repository
	- RepoName: Name of GitHub repository to retrieve issues from
- ZenHub:
	- APIToken: ZenHub API access token  
		    Make an [API token in the ZenHub dashboard](https://app.zenhub.com/dashboard/tokens)

# Undocument ZenHub API Endpoints
ZenHub documents many of their endpoints. However some have not been documented, 
and provide key information.  

## Dependences
Provides information about an issue's dependencies.  

### Request
URL: `api.zenhub.io/v4/repositories/:repo_id/issues/:issue_id/dependencies`  

| Key        | Description                                     |
| ---------- | ----------------------------------------------- |
| `repo_id`  | ID of GitHub repository to retrieve issues from | 
| `issue_id` | ID of GitHub issue to retrieve dependencies for |

### Response
JSON payload with fields:

- `blocked_by` (`User[]`): Array of GitHub issue models blocked by issue
- `blocking` (`User[]`): Array of GitHub issue models blocking current issue

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
				"_id":"5a95d4c9fd6a3f61bd502deb"
			},
			"blocked":true
		}
	],
	"blocking":[]
}
```
