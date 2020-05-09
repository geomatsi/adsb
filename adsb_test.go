// go test -v github.com/skypies/adsb
package adsb

import (
	"testing"
)

func TestAdsbMsgGroundSpeed(t *testing.T) {
	m := Msg{}

	if m.HasGroundSpeed() {
		t.Errorf("ADS-B message should have no ground speed")
	}

	m.SetHasGroundSpeed()

	if !m.HasGroundSpeed() {
		t.Errorf("ADS-B message should have ground speed set")
	}
}

func TestAdsbTrack(t *testing.T) {
	m := Msg{}

	if m.HasTrack() {
		t.Errorf("ADS-B message should have no track")
	}

	m.SetHasTrack()

	if !m.HasTrack() {
		t.Errorf("ADS-B message should have track set")
	}
}

func TestAdsbPosition(t *testing.T) {
	m := Msg{}

	if m.HasPosition() {
		t.Errorf("ADS-B message should have no position")
	}

	m.SetHasPosition()

	if !m.HasPosition() {
		t.Errorf("ADS-B message should have position set")
	}
}
