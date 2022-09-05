package web

type FindByUserIdNoteRequest struct {
	UserId string `json:"id"`
	Limit  int    `json:"limit"`
}
