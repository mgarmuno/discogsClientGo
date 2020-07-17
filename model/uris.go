package model

type Uris struct {
	ID       int    `sql:"integer primary key"`
	ArtistID string `sql:"integer" json:"id"`
	URI      string `sql:"text" json:""`
}
