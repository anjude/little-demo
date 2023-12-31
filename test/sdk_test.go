package test

import (
	"github.com/anjude/bean-sdk-go/beansdk"
	"github.com/anjude/bean-sdk-go/services/csdn_service"
	"testing"
)

func TestSDK(t *testing.T) {
	// visitor token有限
	client := beansdk.NewClient("visitor", "")
	csdnService := csdn_service.NewCsdnService(client, "user name", "user token")
	resp, err := csdnService.HotArticleComment()
	if err != nil {
		t.Errorf("got err: %v", err)
		return
	}
	t.Logf("success: %v", resp)
}
