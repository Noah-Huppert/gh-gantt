import Vue from "vue"
import Vuex from "vuex"
import VuexPersistence from "vuex-persist"

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})

export default new Vuex.Store({
	state: {
		authToken: undefined,
		repositories: undefined,
		pages: {
			home: {
				selectedOwner: "Loading...",
				selectedRepository: "Loading..."
			}
		}
	},
	mutations: {
		authToken(state, authToken) {
			state.authToken = authToken
		},
		repositories(state, repositories) {
			state.repositories = repositories
		},
		homeSelectedOwner(state, selectedOwner) {
			state.pages.home.selectedOwner = selectedOwner
		},
		homeSelectedRepository(state, selectedRepository) {
			state.pages.home.selectedRepository = selectedRepository
		}
	},
	plugins: [vuexLocal.plugin]
})
