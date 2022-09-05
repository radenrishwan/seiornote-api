package domain

type Note struct {
	Id        string
	UserId    string
	Title     string
	Content   string
	Favorite  bool
	CreatedAt string
	UpdatedAt string
}
