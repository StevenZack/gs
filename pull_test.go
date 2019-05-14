package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singlePull(t *testing.T) {
	b, e := singlePull(".")
	assert.NoError(t, e, "msgAndArgs")
	assert.Equal(t, b, false, "msgAndArgs")
}
