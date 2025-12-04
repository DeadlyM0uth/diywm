package main

import (
	"log"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

type Config struct {
	modifier string
	actions  []action
}

type action struct {
	command string
	key     string
}

func main() {

	config := Config{
		modifier: "Mod1",
		actions: []action{
			{
				command: "xclock",
				key:     "c",
			},
			{
				command: "xterm",
				key:     "t",
			},
		},
	}

	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Error establishing connection to X server: %s", err)
	}
	defer conn.Close()

	connInfo := xproto.Setup(conn)
	if connInfo == nil {
		log.Fatal("Couldn't parse connection info")
	}

	root := connInfo.DefaultScreen(conn).Root

	cookie := xproto.ChangeWindowAttributesChecked(
		conn,
		root,
		xproto.CwEventMask,
		[]uint32{
			xproto.EventMaskStructureNotify |
				xproto.EventMaskSubstructureRedirect,
		},
	)
	if cookie.Check() != nil {
		log.Fatal("Is another winodow manager is running?")
	}

	_ = config

	for {

	}
}
