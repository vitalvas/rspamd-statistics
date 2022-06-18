package main

import (
	"log"

	"github.com/vitalvas/rspamd-statistics/app"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	app.NewApp().Execute()
}
