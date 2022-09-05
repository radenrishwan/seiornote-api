package web

type UpdateNoteRequest struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Favorite bool   `json:"favorite"`
}
