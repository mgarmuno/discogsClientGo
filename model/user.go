package model

// User user info
type User struct {
	ID           int    `sql:"integer primary key"`
	DiscogsID    int    `sql:"integer" json:"id"`
	UserName     string `sql:"text" json:"username"`
	ProfileURL   string `sql:"text" json:"resource_url"`
	ConsumerName string `sql:"text" json:"consumer_name"`
}
