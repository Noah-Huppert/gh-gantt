import React from "react"
import PropTypes from "prop-types"

import { HashRouter as Router, Route, Link } from "react-router-dom"

import GitHubLoginContainer from "./ghLogin"

const App = () => (
	<Router>
		<div>
			<a href="/api/v0/auth/login">Login With GitHub</a>
			<Route path="/auth/github" component={GitHubLoginContainer} />
		</div>
	</Router>
)

export default App
