package auth

// AuthToken is a GH Gantt API authentication token
type AuthToken struct {
	// Issuer is the name of the service which issued the token
	Issuer string

	// Subject is the user who the auth token was issued for
	Subject string

	// Audience is the name of the service who is meant to received the token
	Audience string
}
