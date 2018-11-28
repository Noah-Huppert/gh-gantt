import Vue from "vue"
import components from "./components"
import Vue2Storage from "vue2-storage"

Vue.use(Vue2Storage, {
	prefix: "",
	driver: "local",
	ttl: 60 * 60 * 24 * 365 * 100 // 60 seconds in minute -> 60 minutes in hour -> 24 hours in day -> 365 days in year
	// -> 100 years
})

const app = new Vue({
	el: "#app",
	data() {
		return {
			store: this.$storage.get("store", {
				foo: ""
			})
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
	components
})
