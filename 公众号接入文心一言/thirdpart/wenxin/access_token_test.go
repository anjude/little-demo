package wenxin

import (
	"testing"
)

func TestRefreshAccessToken(t *testing.T) {
	RefreshAccessToken()
}

func TestRefreshAccessTokenWithRetry(t *testing.T) {
	got, err := RefreshAccessTokenWithRetry()
	if err != nil {
		t.Errorf("RefreshAccessTokenWithRetry() error = %v", err)
		return
	}
	t.Logf("RefreshAccessTokenWithRetry() got = %v", got)
}
