package storage

import (
	"reflect"
	"testing"

	"ozon/internal/model"
)

func TestCreateComment(t *testing.T) {
	type args struct {
		postID  int
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Comment
		wantErr bool
	}{
		{
			name: "Valid comment",
			args: args{
				postID:  1,
				content: "Some comment",
			},
			want: model.Comment{
				ID:      1,
				PostID:  1,
				Content: "Some comment",
			},
			wantErr: false,
		},
		{
			name: "Comment exceeding length",
			args: args{
				postID:  1,
				content: string(make([]byte, 501)),
			},
			want:    model.Comment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commentIDCounter = 1
			posts = []model.Post{{ID: 1, Title: "Test Post", Content: "Content", Comments: []model.Comment{}}}
			got, err := CreateComment(tt.args.postID, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	type args struct {
		title   string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Post
		wantErr bool
	}{
		{
			name: "Valid post",
			args: args{
				title:   "Tittle post 1",
				content: "Some comment in post 1",
			},
			want: model.Post{
				ID:       1,
				Title:    "Tittle post 1",
				Content:  "Some comment in post 1",
				Comments: []model.Comment{},
			},
			wantErr: false,
		},
		{
			name: "Post content exceeding length",
			args: args{
				title:   "Tittle post",
				content: string(make([]byte, 1001)),
			},
			want:    model.Post{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postIDCounter = 1
			posts = nil
			got, err := CreatePost(tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllPosts(t *testing.T) {
	tests := []struct {
		name string
		want []model.Post
	}{
		{
			name: "Empty posts",
			want: []model.Post{},
		},
		{
			name: "Multi posts",
			want: []model.Post{
				{ID: 1, Title: "Post 1", Content: "Comment 1"},
				{ID: 2, Title: "Post 2", Content: "Comment 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Empty posts" {
				posts = []model.Post{}
			} else {
				posts = []model.Post{
					{ID: 1, Title: "Post 1", Content: "Comment 1"},
					{ID: 2, Title: "Post 2", Content: "Comment 2"},
				}
			}
			if got := GetAllPosts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
