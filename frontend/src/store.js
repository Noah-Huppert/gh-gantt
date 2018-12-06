import Vue from "vue"
import Vuex from "vuex"
import VuexPersistence from "vuex-persist"

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})

export default new Vuex.Store({
	state: {
		authToken: undefined
	},
	mutations: {
		authToken(state, authToken) {
			console.log("commit.authToken", state, authToken)
			state.authToken = authToken
		}
	},
	plugins: [vuexLocal.plugin]
})
