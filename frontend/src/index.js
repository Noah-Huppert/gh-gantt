import Vue from "vue"
import { components, routes } from "./components"
import Vue2Storage from "vue2-storage"
import VueRouter from "vue-router"

// Setup Vue App
// ... Enable developer tools
Vue.config.devtools = true

// ... Store data in local storage
Vue.use(Vue2Storage, {
	prefix: "",
	driver: "local",
	ttl: 60 * 60 * 24 * 365 * 100 // 100 years
})

var store = Vue.$storage.get("store", {
	foo: ""
})

// ... Single page app router
Vue.use(VueRouter)

const router = new VueRouter({
	routes: routes(store)
})

// ... Initialize
const app = new Vue({
	el: "#app",
	data() {
		return {
			store: store	
		}
	},
	watch: {
		store: {
			handler(newStore) {
				this.$storage.set("store", newStore)
			},
			deep: true
		}
	},
	components,
	router
})
