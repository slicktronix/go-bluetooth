// Example use of the btmgmt wrapper
package btmgmt_example

import (
	log "github.com/sirupsen/logrus"
	"github.com/slicktronix/go-bluetooth/hw/linux/btmgmt"
)

func Run() error {

	list, err := btmgmt.GetAdapters()
	if err != nil {
		return err
	}

	for i, a := range list {
		log.Infof("%d) %s (%v)", i+1, a.Name, a.Addr)
	}

	return nil
}
