package web

type CreateNoteRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Favorite bool   `json:"favorite"`
}
