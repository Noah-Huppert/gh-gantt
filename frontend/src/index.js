import Vue from "vue"
import Vuex from "vuex"
import { mapState } from "vuex"
import VueRouter from "vue-router"

import "./sass/styles.sass"
import { components, routes, GHLoginPageRoute, GHLoginCallbackPageRoute, ZenHubLoginPageRoute} from "./components"
import API from "./api"
import store from "./store"
import { parseAuthToken } from "./authToken"

// Setup Vue App
// ... Enable developer tools
Vue.config.debug = true
Vue.config.devtools = true

// ... Store data in local storage

// ... Single page app router
Vue.use(VueRouter)

window.router = new VueRouter({
	routes: routes
})

// ... API client
window.api = new API()

// ... Initialize
const app = new Vue({
	el: "#app",
	computed: mapState([
		"authToken"
	]),
	mounted() {
		// Check if logged into GitHub
		if (this.$router.currentRoute.path != GHLoginCallbackPageRoute && this.authToken === undefined) {

			// If not logged into GitHub redirect to GitHub login page
			router.push(GHLoginPageRoute)
		} else if (parseAuthToken(this.authToken).zenhub_auth_token.length == 0) { // Check if logged into ZenHub
			router.push(ZenHubLoginPageRoute)
		}

	},
	components,
	router,
	store
})
