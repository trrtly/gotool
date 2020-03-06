package lbs

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// LBS struct
type LBS struct {
}

const (
	latCoefficient = 0.009000
	lngCoefficient = 0.010520
)

// RandomLat 纬度
func RandomLat(lat, coverage float64) float64 {
	var min, max float64
	if coverage > 1 {
		min = latCoefficient
		max = latCoefficient * coverage
	} else {
		min = latCoefficient * coverage
		max = latCoefficient
	}

	return format(lat + random(min, max))
}

// RandomLon 经度
func RandomLon(lon, coverage float64) float64 {
	var min, max float64
	if coverage > 1 {
		min = lngCoefficient
		max = lngCoefficient * coverage
	} else {
		min = lngCoefficient * coverage
		max = lngCoefficient
	}
	return format(lon + random(min, max))
}

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func format(num float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", num), 64)
	return value
}
