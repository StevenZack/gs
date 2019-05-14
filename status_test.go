package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gitstatusRepo(t *testing.T) {
	b, e := gitstatusRepo("/Users/stevenzacker/go/src/github.com/StevenZack/progress")
	assert.NoError(t, e, "e")
	assert.Equal(t, b, true, "b")
}
