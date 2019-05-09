package dbops

import (
	"testing"
)

func TestAddUserCredential(t *testing.T) {
	type args struct {
		loginname string
		pwd       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				"heliu",
				"123456",
			},
		},
		{
			name: "case2",
			args: args{
				"chenting",
				"123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUserCredential(tt.args.loginname, tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("AddUserCredential() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		loginname string
		pwd       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				"heliu",
				"123456",
			},
		},
		{
			name: "case2",
			args: args{
				"chenting",
				"123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.loginname, tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserCredential(t *testing.T) {
	type args struct {
		loginname string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{loginname: "heliu"},
			want: "123456",
		},
		{
			name: "case2",
			args: args{loginname: "chenting"},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserCredential(tt.args.loginname)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}
