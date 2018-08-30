package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	now := time.Now()
	ghLoginAttempt := NewGitHubLoginAttempt()

	// Check CreatedOn field is around the current time
	dt := now.Sub(ghLoginAttempt.CreatedOn)
	assert.Equal(t, dt.Seconds() <= 1, true, "GitHubLoginAttempt.CreatedOn is set to a date more than a second from the "+
		"time this test ran")

	// Check State field is 32 characters long
	assert.Equal(t, len(ghLoginAttempt.State), 32, "GitHubLoginAttempt.State must be 32 characters long")
}
