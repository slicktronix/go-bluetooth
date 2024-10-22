package beacon_example

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/slicktronix/go-bluetooth/api/beacon"
)

func Run(beaconType, eddystoneBeaconType, adapterID string) error {

	var b *beacon.Beacon
	if beaconType == "ibeacon" {
		b1, err := beacon.CreateIBeacon("AAAABBBBCCCCDDDDAAAABBBBCCCCDDDD", 111, 999, 89)
		if err != nil {
			return err
		}
		b = b1
	} else {

		if eddystoneBeaconType == "URL" {
			log.Infof("Exposing eddystone URL")
			b1, err := beacon.CreateEddystoneURL("https://bit.ly/2OCrFK2", 99)
			if err != nil {
				return err
			}
			b = b1
		} else {
			// UID
			log.Infof("Exposing eddystone UID")
			b1, err := beacon.CreateEddystoneUID("AAAAAAAAAABBBBBBBBBB", "123456123456", -59)
			if err != nil {
				return err
			}
			b = b1
		}
	}

	// A timeout of 0 cause an immediate timeout and advertisement deregistration
	// see https://www.spinics.net/lists/linux-bluetooth/msg79915.html
	// In seconds
	timeout := uint16(60 * 60 * 18)

	cancel, err := b.Expose(adapterID, timeout)
	if err != nil {
		return err
	}

	defer cancel()

	log.Debugf("%s ready", beaconType)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		os.Exit(0)
	}()

	select {}
}
