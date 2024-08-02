package service

// Post
type Post struct {
	Title    string
	Content  string
	Comments []Comment
}

func (p *Post) AddComment(comment Comment) {
	p.Comments = append(p.Comments, comment)
}
