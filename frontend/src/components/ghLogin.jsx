import React from "react"
import PropTypes from "prop-types"
import { connect } from "react-redux"

/*
 * Indicates what stage of authentication user is currently in.
 * If the stage is empty REDIRECT_FROM_GH is assumed.
 *
 * Flows from one stage to another:
 *
 *	REDIRECT_FROM_GH
 *    (check if `state` and `code` query params are present)
 *        | [if present]      | [if not present]
 *        |                   |
 *        |                   >> BAD_REDIRECT_FROM_GH
 *        v
 *    (try exchanging `state` and `code` for auth token with GH Gantt API)
 *        | [if success]      | [if failure]
 *        |                   |
 *        v                   >> FAILED_SERVER_EXCHANGE
 *  GOOD_SERVER_EXCHANGE
 *
 */
const GitHubAuthStage = {
	REDIRECTED_FROM_GH: "REDIRECTED_FROM_GH", // User was just redirected from GitHub
	BAD_REDIRECT_FROM_GH: "BAD_REDIRECT_FROM_GH", // GitHub redirect was malformed
	FAILED_SERVER_EXCHANGE: "FAILED_SERVER_EXCHANGE", // Failed to exchange GitHub code for auth token
	GOOD_SERVER_EXCHANGE: "GOOD_SERVER_EXCHANGE" // Successfully exchanged GitHub code for auth token
}

/* Component */
class GitHubLogin extends React.Component {
	componentDidMount() {
		if (this.stage === REDIRECT_FROM_GH || this.stage === undefined) {
			// Just redirected from GitHub login
			// Check for `state` and `code` query params
			let params = new URLSearchParams(location.search)

			if (params.get("state") === null || params.get("code") === null) {
				// TODO: Dispatch BAD_REDIRECT_FROM_GH
			}
		}
	}

	render() {
		return (
			<h1>Logging Into GitHub</h1>
		)
	}	
}

GitHubLogin.propTypes = {
	stage: PropTypes.string
}

export { GitHubLogin }

/* Container */
const mapStateToProps = state => {
	return {
		stage: state.auth.gh.stage
	}
}

const mapDispatchToProps = dispatch => {
	return {}
}

const GitHubLoginContainer = connect(
	mapStateToProps,
	mapDispatchToProps
)(GitHubLogin)

export default GitHubLoginContainer
