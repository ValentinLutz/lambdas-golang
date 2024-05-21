package core

import (
	"errors"
	"fmt"

	"github.com/oklog/ulid/v2"
)

type Region string

const (
	RegionEu Region = "EU"
	RegionUs Region = "US"
	RegionCn Region = "CN"
)

var (
	ErrInvalidRegion = errors.New("invalid region")
)

func NewRegion(value string) (Region, error) {
	region := Region(value)
	switch region {
	case RegionEu, RegionUs, RegionCn:
		return region, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrInvalidRegion, value)
	}
}

type OrderId string

func NewOrderId(region Region) OrderId {
	ulidString := ulid.Make().String()
	regionIdentifier := fmt.Sprintf("-%s-", region)
	uildHalfLength := len(ulidString) / 2

	return OrderId(ulidString[0:uildHalfLength] + regionIdentifier + ulidString[uildHalfLength:])
}
