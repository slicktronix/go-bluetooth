package adapter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDiscovery(t *testing.T) {
	a := getDefaultAdapter(t)

	err := a.StartDiscovery()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := a.StopDiscovery()
		if err != nil {
			t.Error(err)
		}
	}()

	discovery, cancel, err := a.OnDeviceDiscovered()
	if err != nil {
		t.Fatal(err)
	}

	var (
		wait = make(chan error, 2)
		wg   sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for dev := range discovery {
			if dev == nil {
				return
			}
			wait <- nil
		}
	}()

	go func() {
		defer wg.Done()
		sleep := 5
		time.Sleep(time.Duration(sleep) * time.Second)
		wait <- fmt.Errorf("Discovery timeout exceeded (%ds)", sleep)
	}()

	err = <-wait
	cancel()

	if err != nil {
		t.Fatal(err)
	}

	wg.Wait()
}

func TestOnDiscoverDataRace(t *testing.T) {
	a, err := GetAdapter(GetDefaultAdapterID())
	if err != nil {
		t.Fatal(err)
	}

	if err = a.FlushDevices(); err != nil {
		t.Fatal(err)
	}

	err = a.StartDiscovery()
	if err != nil {
		t.Fatal(err)
	}

	ch, discoveryCancel, err := a.OnDeviceDiscovered()
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for range ch {

		}
	}()

	cancel := func() {
		err := a.StopDiscovery()
		if err != nil {
			t.Fatal(err)
		}
		discoveryCancel()
	}
	defer cancel()
}
