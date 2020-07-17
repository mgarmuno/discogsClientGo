package model

// Vinyl struct with vinyls informaition
type Vinyl struct {
	ID        int    `sql:"integer primary key"`
	DiscogsID int    `sql:"integer" json:"id"`
	Title     string `sql:"text" json:"title"`
}
