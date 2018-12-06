import { mapState } from "vuex"

import { HomePageRoute } from "."

export default {
	data() {
		return {
			zenhubAuthToken: "",
			authOK: true
		}
	},
	computed: mapState([
		"authToken"
	]),
	template: `
	<div>
		<div v-if="authOK">
			<h1>Enter Your ZenHub Authentication Token</h1>
			<input type="text" v-model="zenhubAuthToken" />
			<button v-on:click="submit">Submit</button>
		</div>

		<div v-if="!authOK">
			<h1>An Error Occurred</h1>
		</div>
	</div>
	`,
	methods: {
		submit() {
			var self = this;

			api.authZenHubAppend(this.authToken, this.zenhubAuthToken)
				.then(authToken => {
					self.$store.commit("authToken", authToken)

					router.push(HomePageRoute)
					window.location.search = ""
				})
				.catch(err => {
					console.error(err)
					self.authOK = false
				})
		}
	}
}
