package service

import (
	"fmt"
	"io"
)

// Forum
type Forum struct {
	writer    io.Writer
	waterball *Waterball
	posts     []Post
	observers []INewPostObserver
}

func (f *Forum) Post(member IMember, post Post) {
	f.posts = append(f.posts, post)

	_, _ = fmt.Fprint(f.writer, fmt.Sprintf("%s: [%s] %s\f", member.Id(), post.Title, post.Content))

	f.Notify(NewPostEvent{PostId: len(f.posts)})
}

func (f *Forum) Comment(postId int, comment Comment) {
	f.GetPost(postId).AddComment(comment)

	_, _ = fmt.Fprint(f.writer, fmt.Sprintf("%s comment in post %d: %s\n", comment.Member.Id(), postId, comment.Content))
}

func (f *Forum) Notify(event NewPostEvent) {
	for _, observer := range f.observers {
		observer.Update(event)
	}
}

func (f *Forum) Register(observer INewPostObserver) {
	f.observers = append(f.observers, observer)
}

func (f *Forum) GetPost(id int) *Post {
	return &f.posts[id-1]
}
