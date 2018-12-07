import { mapState } from "vuex"

export default {
	name: "repo-selector",
	computed: Object.assign({}, mapState([
		"authToken",
		"repositories"
	]), {
		owners() {
			if (this.repositories === undefined) {
				return [ "Loading..." ]
			} else {
				return Object.keys(this.repositories)
			}
		},
		ownerRepositories() {
			if (this.repositories === undefined) {
				return [ "Loading..." ]
			} else {
				return this.repositories[this.selectedOwner]
			}
		}
	}),
	data() {
		return {
			selectedOwner: this.$store.state.pages.home.selectedOwner,
			selectedRepository: this.$store.state.pages.home.selectedRepository
		}
	},
	watch: {
		selectedOwner(selectedOwner) {
			this.$store.commit("homeSelectedOwner", selectedOwner)
		},
		selectedRepository(selectedRepository) {
			this.$store.commit("homeSelectedRepository", selectedRepository)
		}
	},
	template: `
	<div>
		<div class="field repo-select">
			<label class="label">Owner</label>
			<div class="control">
				<div class="select">
					<select v-model="selectedOwner">
						<option v-for="owner in owners">{{ owner }}</option>
					</select>
				</div>
			</div>
		</div>

		<div class="field repo-select">
			<label class="label">Repository</label>
			<div class="control">
				<div class="select">
					<select v-model="selectedRepository">
						<option v-for="repo in ownerRepositories">{{ repo }}</option>
					</select>
				</div>
			</div>
		</div>

	</div>
	`,
	mounted() {
		var self = this

		if (this.repositories === undefined) {
			api.getRepositories(this.authToken)
				.then(repositories => {
					self.$store.commit("repositories", repositories)

					self.selectedOwner = Object.keys(repositories)[0]
					self.selectedRepository = repositories[self.selectedOwner][0]
				})
				.catch(err => {
					console.error(err)
				})
		}
	}
}
