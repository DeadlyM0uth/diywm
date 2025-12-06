package main

import (
	"log"
	"os/exec"

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
		
		// apparently X errors can be ignored????
		ev, xerr := conn.WaitForEvent()
		if ev == nil && xerr == nil {
			log.Fatal("Event and error are nil. Exiting...")
		}

		if ev != nil {
			log.Printf("Event %s\n", ev)
		}
		if xerr != nil {
			log.Printf("Error: %s\n", xerr)
		}
		
		
		switch ev := ev.(type) {
		case xproto.KeyPressEvent:
			//TODO make separate function to handle key presses
			log.Printf("Key pressed: %d\n", ev.Detail)
			// 28 = t
			if ev.Detail == 28 {
				cmd := exec.Command("urxvt")
				cmd.Start()
			}

		case xproto.MapRequestEvent:
			log.Printf("Window wants to be shown: %s", ev.String())

			cookie := xproto.GetWindowAttributes(
				conn,
				ev.Window,
			)

			if winattrib, err := cookie.Reply(); err != nil || !winattrib.OverrideRedirect {
				xproto.MapWindowChecked(conn, ev.Window)
				//TODO make separate function to handle window mapping
				err := xproto.ConfigureWindowChecked(
					conn,
					ev.Window,
					xproto.ConfigWindowX |
					xproto.ConfigWindowY |
					xproto.ConfigWindowWidth |
					xproto.ConfigWindowHeight |
					xproto.ConfigWindowBorderWidth,
					[]uint32{
						0,
						0,
						512,
						512,
						20,
					},
				).Check()
				if err != nil {
					log.Fatal("Error: %w", err)
				}
			}
		}
		
	}
}
