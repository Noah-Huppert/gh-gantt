import { createStore } from "redux"
import { AUTH_TOKEN } from "./actionTypes"

function reducer(state = {}, action) {
	switch (action.type) {
		case AUTH_TOKEN:
			return {
				...state,
				authToken: action.authToken
			}
		default:
			return state
	}
}

export let store = createStore(reducer)
