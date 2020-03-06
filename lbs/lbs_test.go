package lbs

import "testing"

func TestRandomLat(t *testing.T) {
	d := RandomLat(30.05227, 3)
	t.Log(d)
}

func TestRandomLon(t *testing.T) {
	d := RandomLon(121.16765, 5)
	t.Log(d)
}
