package web

import "seiornote/model/domain"

type NoteResponse struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Favorite  bool   `json:"favorite"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewNoteResponse(note domain.Note) NoteResponse {
	return NoteResponse{
		Id:        note.Id,
		UserId:    note.UserId,
		Title:     note.Title,
		Content:   note.Content,
		Favorite:  note.Favorite,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func NewNoteResponses(note []domain.Note) []NoteResponse {
	var notes []NoteResponse

	for _, dummy := range note {
		var note = NoteResponse{
			Id:        dummy.Id,
			UserId:    dummy.UserId,
			Title:     dummy.Title,
			Content:   dummy.Content,
			Favorite:  dummy.Favorite,
			CreatedAt: dummy.CreatedAt,
			UpdatedAt: dummy.UpdatedAt,
		}

		notes = append(notes, note)
	}

	return notes
}
