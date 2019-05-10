package dbops

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
	"video_server/api/defs"
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

func TestAddVideo(t *testing.T) {
	type args struct {
		authorid int
		name     string
	}
	tests := []struct {
		name    string
		args    args
		want    *defs.VideoInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "testcase1",
			args: args{1, "myvideo"},
		},
		{
			name: "testcase2",
			args: args{2, "myvideo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddVideo(tt.args.authorid, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteVideo(t *testing.T) {
	type args struct {
		videoId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "delvideo",
			args: args{
				videoId: "450cd896-e4f7-48ff-85ef-a49af5ca4620",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteVideo(tt.args.videoId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var to int

func TestGetCommentsList(t *testing.T) {
	to, _ = strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	type args struct {
		videoId string
		from    int
		to      int
	}

	tests := []struct {
		name    string
		args    args
		want    []*defs.CommentInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "getcomments",
			args: args{
				"b33afad4-9d3f-4cd9-83ce-b3222d0e7613",
				1514764800,
				to,
			},
			want: []*defs.CommentInfo{
				&defs.CommentInfo{
					"59a2985d-2a19-4ba6-a8b2-5f373805d69e",
					"b33afad4-9d3f-4cd9-83ce-b3222d0e7613",
					"heliu",
					"这是用户1的视频，用户1的评论",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommentsList(tt.args.videoId, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got[0])
				t.Errorf("GetCommentsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddComment(t *testing.T) {
	type args struct {
		videoId  string
		authorId int
		content  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				"b33afad4-9d3f-4cd9-83ce-b3222d0e7613",
				1,
				"这是用户1的视频，用户1的评论",
			},
		},
		{

			args: args{
				"f654556d-847e-4f9d-a667-36891877e820",
				1,
				"这是用户2的视频，用户1的评论",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddComment(tt.args.videoId, tt.args.authorId, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
