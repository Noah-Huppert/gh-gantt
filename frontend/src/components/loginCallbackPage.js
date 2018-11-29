export default {
	props: ["state"],
	data: {
		loginOK: true
	},
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

		// Exchange temporary auth code with server for auth token
		fetch("/api/v0/auth/exchange", {
			method: "POST",
			body: JSON.stringify({
				state: params.get("state"),
				code: params.get("code")
			})
		})
			.then(resp => resp.json())
			.then(resp => {
				self.state.authToken = resp.auth_token
			})
			.catch(err => {
				console.error("Failed to exchange temporary GitHub code with API server", err)
				self.loginOK = false;
			})
	}
}
