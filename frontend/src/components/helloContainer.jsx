import React from "react"
import { connect } from "react-redux"
import { hello } from "../actions"
import Hello from "./hello"

const mapStateToProps = state => {
	return {
		hello: state.hello
	}
}

const mapDispatchToProps = dispatch => {
	return {
		onClick: () => {
			dispatch(hello("world"))
		}
	}
}

const HelloContainer = connect(
	mapStateToProps,
	mapDispatchToProps
)(Hello)

export default HelloContainer
