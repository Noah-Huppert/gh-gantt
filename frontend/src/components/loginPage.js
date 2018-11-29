export default {
	props: [
		"store"
	],
	template: `
	<div>
		<div class="container">
			<a href="/api/v0/auth/login" class="button is-primary">
				Login With GitHub
			</a>
		</div>
	</div>
	`
}
