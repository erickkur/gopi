package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigInstance(t *testing.T) {
	assert := assert.New(t)
	config := GetConfigInstance("database")
	assert.NotEmpty(config.GetString("gopi.engine"))
}
