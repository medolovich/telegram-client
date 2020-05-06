// +build client

package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	envTLGAPIToken = "TLG_API_TOKEN"

	envProxyURL   = "PROXY_URL"
	envProxyLogin = "PROXY_LOGIN"
	envProxyPass  = "PROXY_PASS"
)

var c *TelegramClient

func TestNewTelegramClient(t *testing.T) {
	var err error
	c, err = New(os.Getenv(envTLGAPIToken), os.Getenv(envProxyURL), os.Getenv(envProxyLogin), os.Getenv(envProxyPass))
	assert.Nil(t, err, "expected error to be nil")
	assert.NotNil(t, c, "expected client to be not-nil")
}

func TestGetUpdates(t *testing.T) {
	_, err := c.GetUpdates()
	assert.Nil(t, err, "expected error to be nil")
}
