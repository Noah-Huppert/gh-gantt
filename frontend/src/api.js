/**
 * API Client
 */
export default class API {
	// TODO: Write check resp helper for all api resp
	/**
	 * Make an API request
	 * @param {string} url - URL to make API request to
	 * @param {string} method - HTTP method
	 * @param {object} body - Request body, only valid in non GET requests
	 * @returns {Promise} Resolves with response body object, rejects with an error string
	 */
	makeRequest(url, method, body) {
		var reqOpts = {
			method: method,
		}

		if (method != "GET") {
			reqOpts.body = JSON.stringify(body)
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
		return this.makeRequest("/api/v0/auth/exchange", "POST", {
			state: state,
			code: code
		})
			.then(body => {
				return Promise.resolve(body.auth_token)
			})
			.catch(err => {
				return Promise.reject("error exchanging temporary authentication code with API: " + err)
			})
	}
}
