package service_example

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slicktronix/go-bluetooth/hw"
)

func Run(adapterID string, mode string, hwaddr string) error {

	log.SetLevel(log.TraceLevel)

	btmgmt := hw.NewBtMgmt(adapterID)
	if len(os.Getenv("DOCKER")) > 0 {
		btmgmt.BinPath = "./bin/docker-btmgmt"
	}

	// set LE mode
	btmgmt.SetPowered(false)
	btmgmt.SetLe(true)
	btmgmt.SetBredr(false)
	btmgmt.SetPowered(true)

	if mode == "client" {
		return client(adapterID, hwaddr)
	} else {
		return serve(adapterID)
	}
}
