package libissues

import (
	"fmt"
)

// IssueNode is a node in an issue graph
type IssueNode struct {
	// Issue is the node's issue. If nil the node is considered a root node.
	Issue *Issue

	// NumDescendants holds the number of descendants a node has. If not computed yet then set to -1.
	NumDescendants int

	// Children are the nodes children in the graph
	Children []*IssueNode
}

// NewIssueNode creates an IssueNode
func NewIssueNode(issue *Issue) *IssueNode {
	return &IssueNode{
		Issue:          issue,
		NumDescendants: -1,
		Children:       []*IssueNode{},
	}
}

// String returns a string representation of a graph
func (n IssueNode) String() string {
	str := "["
	if n.Issue != nil {
		str += fmt.Sprintf("%d", n.Issue.Number)
	} else {
		str += "nil"
	}

	str += fmt.Sprintf("(%d)", len(n.Children))

	str += "<"

	for _, child := range n.Children {
		str += child.String()
	}

	str += ">"

	str += "]"

	return str
}

// BuildGraph creates a graph from issues.
// Returns a root node which all other nodes are children of.
func BuildGraph(issues map[int64]*Issue) *IssueNode {
	root := NewIssueNode(nil)
	nodes := map[int64]*IssueNode{}

	for _, issue := range issues {
		insertIssue(root, nodes, issues, issue.Number)
	}

	findNumDescendants(root)

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

// findNumDescendants finds the number of descendants for each node in a graph
func findNumDescendants(root *IssueNode) {
	// If already computed
	if root.NumDescendants >= 0 {
		return
	}

	// If no children
	if len(root.Children) == 0 {
		root.NumDescendants = 0
		return
	}

	// If children
	root.NumDescendants += len(root.Children)

	for _, child := range root.Children {
		// Compute descendants for each child
		findNumDescendants(child)

		root.NumDescendants += child.NumDescendants
	}
}
