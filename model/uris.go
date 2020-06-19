package model

type Uris struct {
	ID       int    `sql:"integer primary key"`
	ArtistID string `sql:"integer" rest:"id"`
	URI      string `sql:"text" rest:""`
}
