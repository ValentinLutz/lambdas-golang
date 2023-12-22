package core

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type Region string

const (
	RegionNone Region = "NONE"
	RegionEu   Region = "EU"
	RegionUs   Region = "US"
)

func NewRegion(value string) (Region, error) {
	region := Region(value)
	switch region {
	case RegionNone, RegionEu, RegionUs:
		return region, nil
	default:
		return "", fmt.Errorf("invalid region '%s'", value)
	}
}

type OrderId string

func NewOrderId(region Region) OrderId {
	ulidString := ulid.Make().String()
	regionIdentifier := fmt.Sprintf("-%s-", region)

	uildHalfLength := len(ulidString) / 2
	return OrderId(ulidString[0:uildHalfLength] + regionIdentifier + ulidString[uildHalfLength:])
}
