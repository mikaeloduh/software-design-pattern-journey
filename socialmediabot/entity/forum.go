package entity

import (
	"fmt"
	"io"
)

// Forum
type Forum struct {
	Writer    io.Writer
	Waterball *Waterball
	Posts     []Post
	observers []INewPostObserver
}

func (f *Forum) Post(member IMember, post Post) {
	f.Posts = append(f.Posts, post)

	_, _ = fmt.Fprint(f.Writer, fmt.Sprintf("%s: [%s] %s\n", member.Id(), post.Title, post.Content))

	f.Notify(NewPostEvent{PostId: len(f.Posts)})
}

func (f *Forum) Comment(postId int, comment Comment) {
	f.GetPost(postId).AddComment(comment)

	_, _ = fmt.Fprint(f.Writer, fmt.Sprintf("%s comment in post %d: %s\n", comment.Member.Id(), postId, comment.Content))
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
	return &f.Posts[id-1]
}

// Post
type Post struct {
	Title    string
	Content  string
	Comments []Comment
}

func (p *Post) AddComment(comment Comment) {
	p.Comments = append(p.Comments, comment)
}

// Comment
type Comment struct {
	Content string
	Member  IMember
}
