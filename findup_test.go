package findup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFromRoot(t *testing.T) {
	_, err := FindFrom(".test", "/")
	assert.EqualError(t, err, "File not found")
}

func TestFindFromDir(t *testing.T) {
	path, err := FindFrom("hosts", "/etc")
	assert.Equal(t, path, "/etc/hosts")
	assert.Equal(t, err, nil)
}

func TestFind(t *testing.T) {
	path, err := Find("etc")
	assert.Equal(t, path, "/etc")
	assert.Equal(t, err, nil)
}
