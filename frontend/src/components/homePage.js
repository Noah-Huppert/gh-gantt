import { GoogleCharts } from "google-charts"
import { mapState } from "vuex"

import { parseAuthToken } from "../authToken"
import { GHLoginPageRoute, GHLoginCallbackPageRoute, ZenHubLoginPageRoute } from "."

export default {
	data() {
		return {
			showGantt: false,
			issues: []
		}
	},
	computed: Object.assign(
		mapState([
			"authToken"
		]), {
		ganttData() {
			if (!this.showGantt) {
				return []
			} else {
				
			}
		},
		ganttOptions() {
			return {}
		}
	}),
	template: `
	<div>
		<div v-if="showGantt">
			<div id="chart"></div>
		</div>

		<div v-if="!showGantt">
			Logging In...
		</div>
	</div>
	`,
	mounted() {
		var self = this

		// Check if logged into GitHub
		if (this.$router.currentRoute.path != GHLoginCallbackPageRoute && 
			this.authToken === undefined) {
			// If not logged into GitHub redirect to GitHub login page
			router.push(GHLoginPageRoute)
		} else if (this.$router.currenRoute != ZenHubLoginPageRoute &&
			parseAuthToken(this.authToken).zenhub_auth_token.length == 0) { // Check if logged into ZenHub
			router.push(ZenHubLoginPageRoute)
		} else { // Fully logged in
			api.getIssues(this.authToken, "Noah-Huppert", "gh-gantt")
				.then(issues => {
					self.issues = issues
					self.showGantt = true

					GoogleCharts.load(() => {
						var rows = [["Issue Number", "Issue Name", "Start Date", "Completion"]]

						for (var i in this.issues) {
							var issue = this.issues[i]
							rows.push([
								issue.number + "",
								issue.title,
								new Date(issue.created_at),
								0.5
							])
						}

						const data = GoogleCharts.api.visualization.arrayToDataTable(rows)

						const chart = new GoogleCharts.api.visualization.Gantt(document.getElementById("chart"))
						chart.draw(data)
						console.log("draw")
					}, {"packages": ["gantt"]})
				})
				.catch(err => {
					console.error(err)
				})
		}
	}
}
