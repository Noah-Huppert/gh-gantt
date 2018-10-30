import { createStore } from "redux"
import { GH_AUTH_REDIRECTED,
	GH_AUTH_BAD_REDIRECT,
	GH_AUTH_BAD_TOKEN_EXCHANGE,
	GH_AUTH_TOKEN_EXCHANGE } from "./actionTypes"

const defaultState = {
	auth: {
		authToken: undefined,
		gh: {
			stage: undefined
		}
	}
}

function reducer(state = defaultState, action) {
	switch (action.type) {
		case GH_AUTH_REDIRECTED:
			state.auth.gh.stage = GH_AUTH_REDIRECTED

			return state

		case GH_AUTH_BAD_REDIRECT:
			state.auth.gh.stage = GH_AUTH_BAD_REDIRECT

			return state

		case GH_AUTH_BAD_TOKEN_EXCHANGE:
			state.auth.gh.stage = GH_AUTH_BAD_TOKEN_EXCHANGE

		case GH_AUTH_TOKEN_EXCHANGE:
			state.auth.authToken = action.authToken

			return state
		default:
			return state
	}
}

export let store = createStore(reducer)
