package wenxin

import (
	"official-account-dev/conf"
	"testing"
)

func TestGetWenXinReply(t *testing.T) {
	conf.InitConf()
	type args struct {
		msgs []Message
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{[]Message{{Role: RoleUser, Content: "你好"}}}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWenXinReply(tt.args.msgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWenXinReply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWenXinReply() got = %v, want %v", got, tt.want)
			}
		})
	}
}
