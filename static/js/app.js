google.charts.load('current', {'packages':['gantt']});
google.charts.setOnLoadCallback(drawChart);

function daysToMilliseconds(days) {
	return days * 24 * 60 * 60 * 1000;
}


// Get data
function drawChart() {
	var req = new XMLHttpRequest();
	req.onreadystatechange = function() {
		if (req.readyState == 4) {
			// Response
			if (req.status !== 200) {
				console.error("error retrieving issue data, http code: "+ 
					req.status + ", response: "+ req.responseText);
			}

			var issues = JSON.parse(req.responseText).issues;

			// Chart
			var data = new google.visualization.DataTable();
			data.addColumn('string', 'Task ID');
			data.addColumn('string', 'Task Name');
			data.addColumn('date', 'Start Date');
			data.addColumn('date', 'End Date');
			data.addColumn('number', 'Duration');
			data.addColumn('number', 'Percent Complete');
			data.addColumn('string', 'Dependencies');


			var rows = []
			for (var i = 0; i < issues.length; i++) {
				var issue = issues[i];

				var deps = null;

				if (issue.blocked_by.length > 0) {
					deps = issue.blocked_by[0].toString();//.join(",");
				}

				var start = null;
				var end = null;

				if (issue.milestone !== undefined) {
					end = new Date(Date.parse(issue.milestone.due_on));
					start = new Date(Date.parse(issue.milestone.due_on));
					start.setDate(start.getDate() - 7);
				} else {
					start = new Date();

					end = new Date();
					end.setDate(end.getDate()+1);
				}

				rows.push([
					issue.number.toString(),
					issue.title,
					start,
					end,
					daysToMilliseconds(1), 
					0,
					""//deps
				]);
			}
			data.addRows(rows);
			/*
			var start = new Date();

			var end = new Date();
			end.setDate(end.getDate()+1);


			data.addRows([
				["1", "first", start, end, daysToMilliseconds(1), 0, ""],
				["2", "second", start, end, daysToMilliseconds(1), 0, "1"]
			]);
			*/

			var chart = new google.visualization.Gantt(document.getElementById("chart"))
			chart.draw(data, { height: 2400 });
		}
	}

	req.open("GET", "/api/issues", true);
	req.send(null);	
}
