/*
 * GH_AUTH prefixed actions indicate what stage of GitHub authentication the user is currently in.
 *
 * Flows from one stage to another:
 *
 *	GH_AUTH_REDIRECTED
 *    (check if `state` and `code` query params are present)
 *        | [if present]      | [if not present]
 *        |                   |
 *        |                   >> GH_AUTH_BAD_REDIRECT
 *        v
 *    (try exchanging `state` and `code` for auth token with GH Gantt API)
 *        | [if success]      | [if failure]
 *        |                   |
 *        v                   >> GH_AUTH_BAD_TOKEN_EXCHANGE
 *  GH_AUTH_TOKEN_EXCHANGE
 *
 */
export const GH_AUTH_REDIRECTED = "GH_AUTH_REDIRECTED"
export const GH_AUTH_BAD_REDIRECT = "GH_AUTH_BAD_REDIRECT"
export const GH_AUTH_BAD_TOKEN_EXCHANGE = "GH_AUTH_BAD_TOKEN_EXCHANGE"
export const GH_AUTH_TOKEN_EXCHANGE = "GH_AUTH_TOKEN_EXCHANGE"

export const ghAuthRedirected = () => {
	return {
		type: GH_AUTH_REDIRECTED
	}
}

export const ghAuthBadRedirect = () => {
	return {
		type: GH_AUTH_BAD_REDIRECT
	}
}

export const ghAuthBadTokenExchange = () => {
	return {
		type: GH_AUTH_BAD_TOKEN_EXCHANGE
	}
}

export const ghAuthTokenExchange = authToken => {
	return {
		type: GH_AUTH_TOKEN_EXCHANGE,
		authToken: authToken
	}
}
