package storage

import (
	"fmt"

	"ozon/internal/model"
)

var posts []model.Post
var comments []model.Comment
var postIDCounter int
var commentIDCounter int

func init() {
	posts = make([]model.Post, 0)
	comments = make([]model.Comment, 0)
	postIDCounter = 1
	commentIDCounter = 1
}

func GetAllPosts() []model.Post {
	return posts
}

func CreatePost(title, content string) (model.Post, error) {
	const postLen = 1000

	if len(content) > postLen {
		return model.Post{}, fmt.Errorf(
			fmt.Sprintf(
				"длина текста комментария не должна превышать %v символов.", postLen,
			),
		)
	}

	post := model.Post{ID: postIDCounter, Title: title, Content: content, Comments: []model.Comment{}}
	posts = append(posts, post)

	postIDCounter++

	return post, nil
}

func CreateComment(postID int, content string) (model.Comment, error) {
	const commentLen = 500

	if len(content) > commentLen {
		return model.Comment{}, fmt.Errorf(
			fmt.Sprintf(
				"длина текста комментария не должна превышать %v символов.", commentLen,
			),
		)
	}

	comment := model.Comment{ID: commentIDCounter, PostID: postID, Content: content}
	comments = append(comments, comment)

	for i, post := range posts {
		if post.ID == postID {
			posts[i].Comments = append(posts[i].Comments, comment)

			break
		}
	}

	commentIDCounter++

	return comment, nil
}
