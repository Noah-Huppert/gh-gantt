import React from "react"
import PropTypes from 'prop-types'

const Hello = ({ onClick, hello }) => (
	<h1 onClick={onClick}>
		i
		Hello {hello}
	</h1>
)

Hello.propTypes = {
	onClick: PropTypes.func.isRequired,
	text: PropTypes.string
}

export default Hello
