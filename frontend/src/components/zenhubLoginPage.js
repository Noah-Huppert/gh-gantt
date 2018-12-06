export default {
	data() {
		return {
			zenhubAuthToken: ""
		}
	},
	template: `
	<div>
		<h1>Enter Your ZenHub Authentication Token</h1>
		<input type="text" v-model="zenhubAuthToken" />
		<button v-on:click="submit">Submit</button>
	</div>
	`,
	methods: {
		submit() {
		}
	}
}
