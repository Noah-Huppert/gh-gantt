package status

// StatusMsg signals the status of a system component
type StatusMsg struct {
	// System is the name of the part of the application the message
	//refers to.
	System string

	// Err holds an error related to a system. This value can be nil to
	// signal that a system succeeded.
	Err error
}
