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
			chartsLoaded: false,
			drawChartOnChartsLoad: false,
			noIssues: false,
			issues: []
		}
	},
	computed: mapState([
			"authToken",
			"pages"
	]),
	components: {
		RepoSelector
	},
	watch: {
		pages: {
			handler() {
				this.loadIssues()
			},
			deep: true
		}
	},
	template: `
	<div>
		<div v-if="showGantt">
			<repo-selector></repo-selector>
			<h1 class="title" v-if="noIssues">No issues</h1>
			<div id="chart" v-if="!noIssues"></div>
		</div>

		<div v-if="!showGantt">
			Logging In...
		</div>
	</div>
	`,
	mounted() {
		// Check if logged into GitHub
		if (this.$router.currentRoute.path != GHLoginCallbackPageRoute && 
			this.authToken === undefined) {
			// If not logged into GitHub redirect to GitHub login page
			router.push(GHLoginPageRoute)
		} else if (this.$router.currenRoute != ZenHubLoginPageRoute &&
			parseAuthToken(this.authToken).zenhub_auth_token.length == 0) { // Check if logged into ZenHub
			router.push(ZenHubLoginPageRoute)
		} else { // Fully logged in
			this.showGantt = true
			GoogleCharts.load(this.onChartsLoaded, {"packages": ["gantt"]})
			this.loadIssues()
		}
	},
	methods: {
		onChartsLoaded() {
			this.chartsLoaded = true
			if(this.drawChartOnChartsLoad) {
				this.drawChart()
			}
		},
		loadIssues() {
			if (this.pages.home.selectedRepository == "Loading...") {
				return
			}

			var self = this

			// Get data
			api.getIssues(this.authToken, this.pages.home.selectedOwner, this.pages.home.selectedRepository)
				.then(issues => {
					self.issues = issues

					if (Object.keys(issues).length == 0) {
						self.noIssues = true
					} else {
						self.noIssues = false
					}


					if (self.chartsLoaded) {
						self.drawChart()
					} else {
						this.drawChartOnChartsLoad = true
					}
				})
				.catch(err => {
					console.error(err)
				})
		},
		drawChart(){
			this.$nextTick(this._drawChart)
		},
		_drawChart() {
			this.drawChartOnChartsLoad = false
			
			if (this.noIssues) {
				return
			}

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

				var deps = null
				if (issue.dependencies !== null) {
					deps = issue.dependencies.join(",")
				}

				rows.push([
					issue.number.toString(), // Issue Number
					issue.title, // Issue Name
					start, // Start Date
					end, // End Date
					24 * 60 * 60 * 1000, // Duration
					0, // Completion
					deps, // Dependencies
				])
			}

			const chart = new GoogleCharts.api.visualization.Gantt(document.getElementById("chart"))	
			const data = GoogleCharts.api.visualization.arrayToDataTable(rows)

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
