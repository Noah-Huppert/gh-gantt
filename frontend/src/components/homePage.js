import { GoogleCharts } from "google-charts"
import { mapState } from "vuex"

import { parseAuthToken } from "../authToken"
import { GHLoginPageRoute, GHLoginCallbackPageRoute, ZenHubLoginPageRoute } from "."
import RepoSelector from "./repoSelector"

export default {
	name: "home-page",
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
	components: {
		RepoSelector
	},
	template: `
	<div>
		<div v-if="showGantt">
			<repo-selector></repo-selector>
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

					GoogleCharts.load(self.drawChart, {"packages": ["gantt"]})
				})
				.catch(err => {
					console.error(err)
				})
		}
	},
	methods: {
		drawChart() {
			// Data format
			// https://developers.google.com/chart/interactive/docs/gallery/ganttchart#data-format
			var rows = [[
				{ label: "Issue Number", type: "string" },
				{ label: "Issue Name", type: "string" },
				{ label: "Start Date", type: "date" },
				{ label: "End Date", type: "date" },
				{ label: "Duration", type: "number" },
				{ label: "Completion", type: "number" },
				{ label: "Dependencies", type: "string" }
			]]

			for (var i in this.issues) {
				var issue = this.issues[i]

				var start = new Date(issue.created_at)

				var end = new Date(issue.created_at)
				end.setDate(end.getDate()+1)

				rows.push([
					issue.number.toString(), // Issue Number
					issue.title, // Issue Name
					start, // Start Date
					end, // End Date
					24 * 60 * 60 * 1000, // Duration
					0, // Completion
					null, // Dependencies
				])
			}

			const data = GoogleCharts.api.visualization.arrayToDataTable(rows)

			const chart = new GoogleCharts.api.visualization.Gantt(document.getElementById("chart"))
			const options = {
				height: 400
			}

			GoogleCharts.api.visualization.events.addListener(chart, "error", (err) => {
				console.error("Gantt Chart error", err)
			})

			chart.draw(data, options)
		}
	}
}
