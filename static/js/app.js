var data = [];

// Get data
var req = new XMLHttpRequest();
req.onreadystatechange = function() {
	if (req.readyState == 4) {
		if (req.status !== 200) {
			console.error("error retrieving issue data, http code: "+ 
				req.status + ", response: "+ req.responseText);
		}

		var issues = JSON.parse(req.responseText).issues;

		for (var i = 0; i < issues.length; i++) {
			var issue = issues[i];

			var deps = null;
			if (issue.blocked_by.length > 0) {
				deps = issue.blocked_by.join(", ");
			}

			data.push({
				id: issue.number.toString(),
				text: issue.title, 
				"start_date":"03-04-2013",
				"duration":"4",
				"progress": 0.1,
				"open": true 
			});
			/*
			{
				id: 1, 
				text: "Project #2", 
				start_date: "01-04-2018", 
				duration: 18, order: 10,
				progress: 0.4, 
				open: true
			}
			*/
		}

		gantt.init("chart");
		gantt.parse(data);
	}
}

req.open("GET", "/api/issues", true);
req.send(null);	
