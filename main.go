package main

import (
	"github.com/mgarmuno/discogsClientGo/db"
	"github.com/mgarmuno/discogsClientGo/gui"
)

func main() {
	db.CheckDatabase()
	gui.InitApp()
}
