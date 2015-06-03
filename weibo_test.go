package weibo

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestClient(t *testing.T) {
	c := &Client{}
	c.ApiKey = "123"
	c.ApiSecret = "456"
	c.RedirectUri = "789"

	authUrl, err := c.GetAuthUrl()

	assert.Equal(t, err, nil)
	assert.Equal(t, authUrl, "https://api.weibo.com/oauth2/authorize?client_id=123&redirect_uri=789&response_type=code")
}
