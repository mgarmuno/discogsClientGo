package model

// Vinyl struct with vinyls informaition
type Vinyl struct {
	ID    int    `sql:"integer primary key"`
	Title string `sql:"text"`
}
