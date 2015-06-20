package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseShipFile(t *testing.T) {
	expected := []string{
		"echo \"shipping!\"",
		"sleep 1",
		"echo \"done shipping\"",
	}
	outcome := parseShipFile()
	assert.Equal(t, expected, outcome)
}
