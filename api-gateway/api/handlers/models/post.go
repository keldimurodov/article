package models

type Post struct {
	Id        string `json:"id"`
	Picture   string `json:"picture"`
	Title     string `json:"name"`
	Article   string `json:"article"`
	OwnerId   string `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdetedAt string `json:"updeted_at"`
	DeletedAt string `json:"deleted_at"`
}
