package utils

import (
	"reflect"
	"testing"
)

func TestSplitMessageByMaxSize(t *testing.T) {
	type args struct {
		msg     string
		maxSize int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		//should split by max size 5
		{args: args{msg: "12345678910", maxSize: 5}, want: []string{"12345", "67891", "0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitMessageByMaxSize(tt.args.msg, tt.args.maxSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitMessageByMaxSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateInvitationCode(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{size: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateInvitationCode(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateInvitationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.args.size {
				t.Errorf("GenerateInvitationCode() got = %v, want %v", got, tt.want)
			}
			println(got)
		})
	}
}
