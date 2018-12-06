/**
 * Parses an authentication token
 * @param {string} authTokenStr - Encoded JWT
 * @throws {string} If an error occurs while parsing the JWT
 * @returns {object} JWT claims
 */
function parseAuthToken(authTokenStr) {
	// Check in correct format
	let parts = authTokenStr.split(".")

	if (parts.length != 3) {
		throw "JWT must be in format <header>.<claims>.<signature>"
	}

	// Base 64 decode
	let jsonStr = atob(parts[1])

	// Parse JSON
	return JSON.parse(jsonStr)
}
