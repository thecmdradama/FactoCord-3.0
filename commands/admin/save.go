package admin

import (
	"github.com/bwmarrin/discordgo"

	"github.com/maxsupermanhd/FactoCord-3.0/support"
)

// SaveServer executes the save command on the server.
func SaveServer(s *discordgo.Session, args string) {
	if len(args) != 0 {
		support.Send(s, "Save accepts no arguments")
		return
	}
	success := support.SendToFactorio("/save")
	if success {
		// TODO read log to be sure it's successful
		support.Send(s, "Server saved successfully!")
	} else {
		support.Send(s, "Sorry, there was an error sending /save command")
	}
}
