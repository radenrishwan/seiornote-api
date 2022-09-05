package web

type FindFavoriteNoteRequest struct {
	UserId string `json:"user_id"`
	Limit  int    `json:"limit"`
}
