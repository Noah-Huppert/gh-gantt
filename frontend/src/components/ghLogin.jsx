import React from "react"
import PropTypes from "prop-types"
import { connect } from "react-redux"

/* Component */
const GitHubLogin = ({ onClick, hello }) => (
	<h1>Logging Into GitHub</h1>
)

GitHubLogin.propTypes = {}

export { GitHubLogin }

/* Container */
const mapStateToProps = state => {
	return {}
}

const mapDispatchToProps = dispatch => {
	return {}
}

const GitHubLoginContainer = connect(
	mapStateToProps,
	mapDispatchToProps
)(GitHubLogin)

export default GitHubLoginContainer
