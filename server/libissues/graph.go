package libissues

// IssueNode is a node in an issue graph
type IssueNode struct {
	// Issue is the node's issue. If nil the node is considered a root node.
	Issue *Issue

	// Children are the nodes children in the graph
	Children []*IssueNode
}

// NewIssueNode creates an IssueNode
func NewIssueNode(issue *Issue) *IssueNode {
	return &IssueNode{
		Issue:    issue,
		Children: []*IssueNode{},
	}
}

// BuildGraph creates a graph from issues.
// Returns a root node which all other nodes are children of.
func BuildGraph(issues map[int64]*Issue) *IssueNode {
	root := NewIssueNode(nil)
	nodes := map[int64]*IssueNode{}

	for _, issue := range issues {
		insertIssue(root, nodes, issues, issue.Number)
	}

	return root
}

// insertIssue is a recursive function which adds an issue and all its dependencies into a graph.
// The inserted argument tracks issues which have already been inserted into the graph.
func insertIssue(root *IssueNode, nodes map[int64]*IssueNode, issues map[int64]*Issue, issueNumber int64) {
	issue := issues[issueNumber]

	// Ensure all dependencies are inserted
	for _, depNumber := range issue.Dependencies {
		// If dependency not inserted
		if _, ok := nodes[depNumber]; !ok {
			// Insert dependency first
			insertIssue(root, nodes, issues, depNumber)
		}
	}

	// Create IssueNode
	nodes[issueNumber] = NewIssueNode(issues[issueNumber])
	node := nodes[issueNumber]

	// Insert into all dependencies or root if it has none
	if len(issue.Dependencies) == 0 {
		root.Children = append(root.Children, node)
	} else {
		for _, depNumber := range issue.Dependencies {
			depNode := nodes[depNumber]
			depNode.Children = append(depNode.Children, node)
			nodes[depNumber] = depNode
		}
	}
}
