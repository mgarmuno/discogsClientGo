package model

// Artist struct with artist information
type Artist struct {
	ID        int    `sql:"integer primary key"`
	DiscogsID int    `sql:"integer" rest:"id"`
	Name      string `sql:"text" rest:"name"`
	Profile   string `sql:"text" rest:"profile"`
	URI       string `sql:"text" rest:"uri"`
}
