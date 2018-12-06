# Run Book
GH Gantt Server run book.

# Table Of Contents
- [Setup GitHub Application](#setup-github-application)

# Setup GitHub Application
A GitHub application must be setup so users can login with GitHub.  

1. Navigate to the [GitHub Developer Settings page](https://github.com/settings/developers)
2. Create a new GitHub application
	- Click the "New OAuth App" button in the top right of the page
	- The values of the "Application name", and "Application description" field are not important
	- Set the "Homepage URL" to the subdomain and host of your GH Gantt server
		- Ex: `gh-gantt.example.com`
	- Set the "Authorization callback URL"
		- Should be the "Homepage URL" value with a fragment portion of `#/auth/github_callback`
		- Ex: `gh-gantt.example.com/#/auth/github_callback`
	- Click the "Register Application" button
3. Save the application's credentials
	- Save your application's "Client ID" and "Client Secret" in the `APP_GITHUB_CLIENT_ID` and
		`APP_GITHUB_CLIENT_SECRET` environment variables
