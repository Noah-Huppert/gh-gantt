import { createStore } from "redux"
import { HELLO } from "./actionTypes"

function reducer(state = {}, action) {
	switch (action.type) {
		case HELLO:
			return {
				...state,
				hello: "world"
			}
		default:
			return state
	}
}

export let store = createStore(reducer)
