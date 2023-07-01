package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidIPv4Address(t *testing.T) {
	assert.True(t, IsValidIPv4Address("192.168.1.1"))
	assert.True(t, IsValidIPv4Address("0.0.0.0"))
	assert.True(t, IsValidIPv4Address("255.255.255.255"))
	assert.False(t, IsValidIPv4Address("255.255.255.256"))
	assert.False(t, IsValidIPv4Address("::1"))
	assert.False(t, IsValidIPv4Address("localhost"))
}
