import React from "react"
import PropTypes from "prop-types"
import { connect } from "react-redux"

import { * } from "../actions"

/* Component */
class GitHubLogin extends React.Component {
	componentDidMount() {
		if (this.stage === undefined) {
			dispatch(ghAuthRedirected())
		} else if (this.stage === GH_AUTH_REDIRECTED) { // Just redirected from GitHub login
			// Check for `state` and `code` query params
			let params = new URLSearchParams(location.search)

			if (params.get("state") === null || params.get("code") === null) {
				dispatch(ghAuthBadRedirect())
				return
			}

			// 
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
