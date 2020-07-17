package model

// Artist struct with artist information
type Artist struct {
	ID        int    `sql:"integer primary key"`
	DiscogsID int    `sql:"integer" json:"id"`
	Name      string `sql:"text" json:"name"`
	Profile   string `sql:"text" json:"profile"`
	URI       string `sql:"text" json:"uri"`
}
