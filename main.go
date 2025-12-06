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

	_ = config

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
			xproto.EventMaskKeyPress |
				xproto.EventMaskStructureNotify |
				xproto.EventMaskSubstructureRedirect,
		},
	)
	if cookie.Check() != nil {
		log.Fatal("Is another winodow manager is running?")
	}

	for {

		ev, xerr := conn.WaitForEvent()
		if ev == nil && xerr == nil {
			log.Fatal("Event and error are nil. Exxiting...")
		}

		if ev != nil {
			log.Printf("Event %s\n", ev)
		}
		if xerr != nil {
			log.Printf("Error: %s\n", xerr)
		}

		switch ev.(type) {
		case xproto.KeyPressEvent:
			kpe := ev.(xproto.KeyPressEvent)
			log.Printf("Key pressed: %d\n", kpe.Detail)
		}

	}
}
