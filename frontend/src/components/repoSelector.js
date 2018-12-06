import { mapState } from "vuex"

export default {
	name: "repo-selector",
	computed: mapState([
		"authToken"
	]),
	template: `
	<div>
	</div>
	`,
	mounted() {
		api.getRepositories(this.authToken)
			.then(repositories => {
				console.log(repositories)
			})
			.catch(err => {
				console.error(err)
			})
	}
}
