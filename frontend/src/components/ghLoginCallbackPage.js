import { mapState } from "vuex"

import { HomePageRoute } from "."

export default {
	name: "gh-login-callback-page",
	data() {
		return {
			loginOK: true
		}
	},
	computed: mapState([
		"authToken"
	]),
	template: `
	<div>
		<div class="container">
			<h1 v-if="loginOK" class="title">Logging In...</h1>

			<h1 v-if="!loginOK" class="title">Login Error</h1>
		</div>
	</div>
	`,
	mounted() {
		var self = this;

		// Check query parameters exist
		const params = new URLSearchParams(window.location.search)

		if (!params.has("state") || !params.has("code")) {
			console.error("URL does not have required query parameters for login")

			this.loginOK = false;
			return;
		}

		const state = params.get("state") 
		const code = params.get("code")

		// Exchange temporary auth code with server for auth toke		
		api.authExchange(state, code)
			.then(authToken => {
				self.$store.commit("authToken", authToken)

				router.push(HomePageRoute)
				window.location.search = ""
			})
			.catch(err => {
				console.error(err)

				self.loginOK = false;
			})
	}
}
