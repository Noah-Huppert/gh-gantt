/**
 * API Client
 */
export default class API {
	/**
	 * Make an API request
	 * @param {string} url - URL to make API request to
	 * @param {string} method - HTTP method
	 * @param {string} authToken - API authentication token, if not needed pass undefined
	 * @param {object} body - Request body, only valid in non GET requests, if not needed pass undefined
	 * @param {object} query - Request query parameters, if not needed pass undefined
	 * @returns {Promise} Resolves with response body object, rejects with an error string
	 */
	makeRequest(url, method, authToken, body, query) {
		var reqOpts = {
			method: method,
			headers: {}
		}

		if (authToken !== undefined) {
			reqOpts.headers["Authorization"] = "token " + authToken
		}

		if (body !== undefined && method != "GET") {
			reqOpts.body = JSON.stringify(body)
		}

		if (query !== undefined) {
			url += "?"

			var first = true
			for (var key in query) {
				if (!first) {
					url += "&"
				}

				url += key + "=" + encodeURIComponent(query[key])

				first = false
			}
		}

		return fetch(url, reqOpts)
			.then(resp => {
				return resp.json()
					.then(body => {
						return Promise.resolve([resp, body])
					})
			})
			.catch(err => {
				return Promise.reject("error decoding response body into JSON: " + err)
			})
			.then(([resp, body]) => {
				// Check response code
				if (resp.status != 200) {
					return Promise.reject("error returned by API: " + body.error)
				}

				// Success
				return Promise.resolve(body)
			})
	}

	/**
	 * Exchanges a temporary GitHub code for an API auth token
	 * @param {string} state - Value returned by GitHub to prevent cross site request forgery
	 * @param {string} code - Short lived GitHub authentication token 
	 * @returns {Promise} Resolves with API authentication token, rejects with error string
	 */
	authExchange(state, code) {
		return this.makeRequest("/api/v0/auth/exchange", "POST", undefined, {
			state: state,
			code: code
		}, undefined)
			.then(body => {
				return Promise.resolve(body.auth_token)
			})
			.catch(err => {
				return Promise.reject("error exchanging temporary authentication code with API: " + err)
			})
	}

	/**
	 * Sends a ZenHub authentication token to the server to be appended to an existing authentication token
	 * @param {string} authToken - Existing API authentication token
	 * @param {string} zenhubAuthToken - ZenHub authentication token
	 * @returns {Promise} Resolves with new API authentication token, rejects with error string
	 */
	authZenHubAppend(authToken, zenhubAuthToken) {
		return this.makeRequest("/api/v0/auth/zenhub", "POST", undefined, {
			auth_token: authToken,
			zenhub_auth_token: zenhubAuthToken
		}, undefined)
			.then(body => {
				return Promise.resolve(body.auth_token)
			})
			.catch(err => {
				return Promise.reject("error appending ZenHub authentication token: " + err)
			})
	}

	/**
	 * Retrieves a list of GitHub issues for a repository
	 * @param {string} authToken - API authentication token
	 * @param {string} owner - Repository owner
	 * @param {stirng} name - Repository name
	 * @returns {Promise} Resolves with array of issues, rejects with error string
	 */
	getIssues(authToken, owner, name) {
		return this.makeRequest("/api/v0/issues", "GET", authToken, undefined, {
			repository_owner: owner,
			repository_name: name
		})
			.then(body => {
				return Promise.resolve(body.issues)
			})
			.catch(err => {
				return Promise.reject("error retrieving GitHub issues from API: " + err)
			})
	}

	/**
	 * Retrieves a list of GitHub repositories
	 * @param {string} authToken - API authentication token
	 * @returns {Promise} Resolves with object containing repositories. Keys are organizations. Values are arrays of 
	 * repository names. Rejects with an error string
	 */
	getRepositories(authToken) {
		return this.makeRequest("/api/v0/repositories", "GET", authToken, undefined, undefined)
			.then(body => {
				return Promise.resolve(body.repositories)
			})
			.catch(err => {
				return Promise.reject("error retrieving GitHub repositories from API: " + err)
			})
	}
}
