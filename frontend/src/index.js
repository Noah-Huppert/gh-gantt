import Vue from "vue"
import Vuex from "vuex"
import { mapState } from "vuex"
import VueRouter from "vue-router"

import "./sass/styles.sass"
import { components, routes, LoginPageRoute, LoginCallbackPageRoute } from "./components"
import API from "./api"
import store from "./store"

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
		// ... Check if logged in
		if (this.$router.currentRoute.path != LoginCallbackPageRoute && this.authToken === undefined) {

			// If not logged in
			router.push(LoginPageRoute)
		}
	},
	components,
	router,
	store
})
